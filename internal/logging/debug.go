package logging

func (l *Log) Debug(s string) {
	if l.LogLevel < debugLevel {
		return
	}
	l.ZeroLogger.Debug().Msg(s)
}
