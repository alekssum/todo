package logging

func (l *Log) Warn(s string) {
	if l.LogLevel < warnLevel {
		return
	}

	l.ZeroLogger.Warn().Msg(s)
}
