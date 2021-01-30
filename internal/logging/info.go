package logging

func (l *Log) Info(s string) {
	if l.LogLevel < infoLevel {
		return
	}

	l.ZeroLogger.Info().Msg(s)
}
