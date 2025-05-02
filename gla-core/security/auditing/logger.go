package auditing

import (
	"fmt"
	"log"
	"os"
)

// Logger is used to handle the logging of audit events.
type Logger struct {
	file *os.File
}

// NewLogger creates a new logger instance.
func NewLogger(logFile string) (*Logger, error) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Logger{file: file}, nil
}

// Log writes an audit log message to the file.
func (l *Logger) Log(message string) {
	log.SetOutput(l.file)
	log.Println(message)
}

// Close closes the log file.
func (l *Logger) Close() {
	l.file.Close()
}
