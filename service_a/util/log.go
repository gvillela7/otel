package util

import (
	"log/slog"
	"os"
)

func logConf(logDefault, logJson bool) *slog.Logger {
	if logDefault {
		return slog.Default()
	} else if logJson {
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	return slog.Default()
}

func Log(logDefault, logJson bool, logType, message string, args ...any) *slog.Logger {
	logger := logConf(logDefault, logJson)
	switch logType {
	case "error":
		logger.Error(message, args...)
	case "info":
		logger.Info(message, args...)
	case "warn":
		logger.Warn(message, args...)
	case "debug":
		logger.Debug(message, args...)
	}
	return logger
}
