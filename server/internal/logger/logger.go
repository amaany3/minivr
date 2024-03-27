package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"

	"cloud.google.com/go/logging"
)

var (
	// levelDefault = slog.Level(logging.Default)
	LevelDebug    = slog.Level(logging.Debug)
	LevelInfo     = slog.Level(logging.Info)
	LevelNotice   = slog.Level(logging.Notice)
	LevelWarning  = slog.Level(logging.Warning)
	LevelError    = slog.Level(logging.Error)
	LevelCritical = slog.Level(logging.Critical)
	// levelAlert     = slog.Level(logging.Alert)
	// levelEmergency = slog.Level(logging.Emergency)
)

func EnableCloudLoggingLogger() {
	l := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.MessageKey:
				a.Key = "message"
			case slog.SourceKey:
				a.Key = "logging.googleapis.com/sourceLocation"
			case slog.LevelKey:
				a.Key = "severity"
				a.Value = slog.StringValue(logging.Severity(a.Value.Any().(slog.Level)).String())
			}
			return a
		},
	})
	slog.SetDefault(slog.New(l))
}

// Debug or trace information
func Debug(format string, args ...any) {
	logAttrs(LevelDebug, format, args...)
}

// Routine information, such as ongoing status or performance
func Info(format string, args ...any) {
	logAttrs(LevelInfo, format, args...)
}

// Normal but significant events, such as start up, shut down, or a configuration change
func Notice(format string, args ...any) {
	logAttrs(LevelNotice, format, args...)
}

// Warning events might cause problems
func Warn(format string, args ...any) {
	logAttrs(LevelWarning, format, args...)
}

// Error events are likely to cause problems
func Error(format string, args ...any) {
	logAttrs(LevelError, format, args...)
}

// Critical events cause more severe problems or outages
func Critical(format string, args ...any) {
	logAttrs(LevelCritical, format, args...)
}

// skip caller for add source
// https://github.com/golang/go/issues/59145#issuecomment-1481920720
func logAttrs(level slog.Level, format string, args ...any) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	message := fmt.Sprintf(format, args...)
	r := slog.NewRecord(time.Now(), level, message, pcs[0])
	_ = slog.Default().Handler().Handle(context.Background(), r)
}
