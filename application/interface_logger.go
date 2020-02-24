package application

type LoggerInterface interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Critical(message string)
}
