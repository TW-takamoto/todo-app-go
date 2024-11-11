package error_test

import (
	"testing"

	e "api/internal/shared/error"
)

func TestNewError(t *testing.T) {
	tests := []struct {
		name string
		args error
		want string
	}{
		{"NotFound", e.ErrorBuilder(e.NotFound).Build(), "データがありませんでした"},
		{"InvalidArgument", e.ErrorBuilder(e.InvalidArgument).Build(), "パラメータが不正です"},
		{"InvalidArgument(ID)", e.ErrorBuilder(e.InvalidArgument).Property("ID").Build(), "IDが不正です"},
		{"Custom", e.ErrorBuilder(e.Custom).Property("Custom").Build(), "Custom"},
		{"Unknown", e.ErrorBuilder(e.Unknown).Build(), "予期せぬエラーが発生しました"},
		{"Database", e.ErrorBuilder(e.Database).Build(), "データベースの処理に失敗しました"},
		{"Unknown(-1)", e.ErrorBuilder(-1).Build(), "予期せぬエラーが発生しました"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Error()
			if got != tt.want {
				t.Errorf("%v is not match: (%v), (%v)", tt.name, got, tt.want)
			}
		})
	}

}
