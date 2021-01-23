package logging

import "fmt"

func (l *Log) Info(v ...interface{}) {
	l.ZeroLogger.Info().Msg(fmt.Sprint(v))
	l.Pipe <- fmt.Sprint(v)
	l.Pipe <- fmt.Sprint(v)
}
