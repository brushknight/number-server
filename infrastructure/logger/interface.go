package logger

type LoggerInterface interface {
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Critical(message string)
}
