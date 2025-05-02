package auditing

import (
	"fmt"
	"time"
)

// Auditor handles audit operations like logging system activities.
type Auditor struct {
	logger *Logger
}

// NewAuditor creates a new Auditor instance.
func NewAuditor(logFile string) (*Auditor, error) {
	logger, err := NewLogger(logFile)
	if err != nil {
		return nil, err
	}
	return &Auditor{logger: logger}, nil
}

// RecordAudit logs a system event with timestamp and description.
func (a *Auditor) RecordAudit(eventDescription string) {
	timestamp := time.Now().Format(time.RFC3339)
	message := fmt.Sprintf("%s - Event: %s", timestamp, eventDescription)
	a.logger.Log(message)
}

// Close closes the logger in the auditor.
func (a *Auditor) Close() {
	a.logger.Close()
}
