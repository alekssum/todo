package logging

func (l *Log) Error(err error) {
	l.ZeroLogger.Err(err).Msg("")
}
