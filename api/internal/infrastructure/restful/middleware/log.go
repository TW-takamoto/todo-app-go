package middleware

import (
	"api/internal/domain/primitive"
	e "api/internal/shared/error"

	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
		// Swaggerはスキップ
		path := c.Request.URL.Path
		if strings.Contains(path, "swagger") {
			c.Next()
			return
		}
		// レスポンスのbodyを出力するためのライター
		blw := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		// リクエストの開始時刻
		start := time.Now()
		// リクエストの識別子
		id_, err := primitive.NewID()
		var id string
		if err != nil {
			id = "00000000-0000-0000-0000-000000000000"
		} else {
			id = id_.Value()
		}
		// リクエストボディを読み込む
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			logger.Error("failed to read request body", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to read request body"})
			return
		}
		// リクエストボディをリセット
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		// リクエストボディをログ出力
		logger.Info("request",
			zap.String("id", id),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.ByteString("body", bodyBytes),
			zap.String("ip", c.ClientIP()),
		)
		// リクエストの処理
		c.Next()
		// レスポンスの処理時間
		duration := time.Since(start)
		// レスポンスのログ出力
		logger.Info("response",
			zap.String("id", id),
			zap.Duration("duration", duration),
			zap.Int("status", c.Writer.Status()),
			zap.String("body", blw.body.String()),
			zap.String("ip", c.ClientIP()),
		)
	}
}

func getLogger(c *gin.Context) (*zap.Logger, error) {
	logger, ok := c.Get("logger")
	if !ok {
		return &zap.Logger{}, e.ErrorBuilder(e.NotFound).Build()
	}
	logger_, ok := logger.(*zap.Logger)
	if !ok {
		return &zap.Logger{}, e.ErrorBuilder(e.Unknown).Build()
	}
	return logger_, nil
}
