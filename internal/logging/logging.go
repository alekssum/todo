package logging

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

const (
	errorLevel = 1
	warnLevel  = 2
	infoLevel  = 3
	debugLevel = 4

	ErrorLevelName = "error"
	WarnLevelName  = "warn"
	InfoLevelName  = "info"
	DebugLevelName = "debug"
)

type Logger interface {
	Info(s string)
	Error(err error)
	Warn(s string)
	Debug(s string)
}

func New(cfg *Config) *Log {
	l := zlog.Logger.With().Caller().Logger()
	return &Log{
		ZeroLogger: &l,
		LogLevel:   levelNumber(cfg.LogLevel),
	}
}

type Log struct {
	ZeroLogger *zerolog.Logger
	LogLevel   int
}

func levelNumber(lvl string) int {
	switch lvl {
	case ErrorLevelName:
		return errorLevel
	case WarnLevelName:
		return warnLevel
	case DebugLevelName:
		return debugLevel
	default:
		return infoLevel
	}
}
