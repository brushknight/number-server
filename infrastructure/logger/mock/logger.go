package mock

type Logger struct {
	LastMessage string
}

func (l *Logger) Trace(message string) {
	l.LastMessage = "[TRACE] " + message
}

func (l *Logger) Debug(message string) {
	l.LastMessage = "[DEBUG] " + message
}
func (l *Logger) Info(message string) {
	l.LastMessage = "[INFO] " + message

}
func (l *Logger) Warn(message string) {
	l.LastMessage = "[WARN] " + message

}
func (l *Logger) Error(message string) {
	l.LastMessage = "[ERROR] " + message

}
func (l *Logger) Critical(message string) {
	l.LastMessage = "[CRITICAL] " + message
}

func NewMockLogger() *Logger {
	return &Logger{}
}
