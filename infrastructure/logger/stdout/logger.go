package stdout

import (
	"fmt"
	"time"
)

type Logger struct {
	env string
}

func (l *Logger) Trace(message string) {
	if l.env == "dev" {
		//logToStdout("[TRACE] " + message)
	}
}

func (l *Logger) Debug(message string) {
	if l.env == "dev" {
		logToStdout("[DEBUG] " + message)
	}
}

func (l *Logger) Info(message string) {
	logToStdout("[INFO] " + message)

}
func (l *Logger) Warn(message string) {
	logToStdout("[WARN] " + message)

}
func (l *Logger) Error(message string) {
	logToStdout("[ERROR] " + message)

}
func (l *Logger) Critical(message string) {
	logToStdout("[CRITICAL] " + message)
	panic(message)
}

func NewLogger(env string) *Logger {
	return &Logger{env: env}
}

func logToStdout(message string) {
	t := time.Now()
	timePrefix := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	fmt.Printf("%s %s\n", timePrefix, message)
}
