package logging

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type Logger interface {
	Log(v ...interface{})
	Info(v ...interface{})
	Error(v ...interface{})
	Warn(v ...interface{})
}

type LogSubscriber interface {
	Send(v ...interface{})
}

func New(cfg *Config) *Log {
	l := zlog.Logger.With().Caller().Logger()
	return &Log{
		ZeroLogger: &l,
		Pipe:       make(chan interface{}),
	}
}

type Log struct {
	ZeroLogger *zerolog.Logger
	Pipe       chan interface{}
}

func (l *Log) Subscribe(sl []LogSubscriber) {

	go func() {
		for v := range l.Pipe {
			for i := range sl {
				sl[i].Send(v)
			}
		}
	}()

}
