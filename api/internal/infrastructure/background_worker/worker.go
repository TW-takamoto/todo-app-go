package background_worker

import (
	"time"

	"go.uber.org/zap"
)

type Worker struct {
	interval               time.Duration
	logger                 *zap.Logger
	execute_for_first_time bool
}

func (w Worker) Execute(callback func() error) error {
	exec := Executer(callback, w.logger)
	if w.execute_for_first_time {
		exec()
	}
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()
	next := time.Now().Add(w.interval)
	time.Sleep(time.Until(next))
	for {
		select {
		case <-ticker.C:
			exec()
		}
	}
}

func Executer(callback func() error, logger *zap.Logger) func() {
	return func() {
		if logger != nil {
			logger.Info("Worker", zap.Time("time", time.Now()))
		}
		err := callback()
		if err != nil {
			if logger != nil {
				logger.Error(err.Error())
			}
		}
	}
}

func NewWorker(interval *time.Duration, execute_for_first_time bool, logger *zap.Logger) Worker {
	interval_ := time.Duration(1) * time.Hour
	if interval != nil {
		interval_ = *interval
	}
	return Worker{interval: interval_, logger: logger, execute_for_first_time: execute_for_first_time}
}
