package log

import (
	"os"
	"time"

	kitlog "github.com/go-kit/kit/log"
)

var log = newLogWrapper()

type logWrapper struct {
	info  *kitlog.Context
	error *kitlog.Context
}

func (wrapper *logWrapper) Info(v ...interface{}) {
	wrapper.info.WithPrefix("time", time.Now().Format(time.RFC3339)).Log(v...)
}

func (wrapper *logWrapper) Error(v ...interface{}) {
	wrapper.error.WithPrefix("time", time.Now().Format(time.RFC3339)).Log(v...)
}

func newLogWrapper() *logWrapper {
	return &logWrapper{
		kitlog.NewContext(kitlog.NewLogfmtLogger(os.Stdout)).WithPrefix("level", "info"),
		kitlog.NewContext(kitlog.NewLogfmtLogger(os.Stderr)).WithPrefix("level", "error"),
	}
}

func Info(v ...interface{}) {
	log.Info(v...)
}

func Error(v ...interface{}) {
	log.Error(v...)
}
