package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

const logID = "log_id"

type DdserLogger struct {
	Hostname string
	*logrus.Logger
}

func MustGetLogger() *DdserLogger {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	DdLogger := &DdserLogger{hostname, logrus.StandardLogger()}
	return DdLogger
}

func (lg *DdserLogger) Info(args ...interface{}) {
	fields(lg).Info(args...)
}

func (lg *DdserLogger) Infof(format string, args ...interface{}) {
	fields(lg).Infof(format, args...)
}

func (lg *DdserLogger) Debug(args ...interface{}) {
	fields(lg).Debug(args...)
}

func (lg *DdserLogger) Debugf(format string, args ...interface{}) {
	fields(lg).Debugf(format, args...)
}

func (lg *DdserLogger) Warn(args ...interface{}) {
	fields(lg).Warn(args...)
}

func (lg *DdserLogger) Warnf(format string, args ...interface{}) {
	fields(lg).Warnf(format, args...)
}

func (lg *DdserLogger) Error(args ...interface{}) {
	fields(lg).Error(args...)
}

func (lg *DdserLogger) Errorf(format string, args ...interface{}) {
	fields(lg).Errorf(format, args...)
}

// DebugfWithId write formatted debug level log with added log_id field
func (lg *DdserLogger) DebugfWithId(id string, format string, args ...interface{}) {
	fields(lg).WithField(logID, id).Debugf(format, args...)
}

// InfofWithId write formatted info level log with added log_id field
func (lg *DdserLogger) InfofWithId(id string, format string, args ...interface{}) {
	fields(lg).WithField(logID, id).Infof(format, args...)
}

// InfoWithId write info level log with added log_id field
func (lg *DdserLogger) InfoWithId(id string, args ...interface{}) {
	fields(lg).WithField(logID, id).Info(args...)
}

// ErrorfWithId write formatted error level log with added log_id field
func (lg *DdserLogger) ErrorfWithId(id string, format string, args ...interface{}) {
	fields(lg).WithField(logID, id).Errorf(format, args...)
}

// ErrorWithId write error level log with added log_id field
func (lg *DdserLogger) ErrorWithId(id string, args ...interface{}) {
	fields(lg).WithField(logID, id).Error(args...)
}

func fields(lg *DdserLogger) *logrus.Entry {
	file, line := getCaller()
	return lg.Logger.WithField("hostname", lg.Hostname).WithField("time", time.Now().UTC().Format(time.RFC3339)).WithField("source", fmt.Sprintf("%s:%d", file, line))
}

func getCaller() (string, int) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	return file, line
}
