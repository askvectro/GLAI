package security

import (
    "fmt"
    "gla-core/security/auditing"
    "gla-core/security/sanitization"
)

// SecurityManager links the sanitization and auditing functionalities.
type SecurityManager struct {
    auditor *auditing.Auditor
}

// NewSecurityManager creates a new instance of SecurityManager.
func NewSecurityManager(logFile string) (*SecurityManager, error) {
    auditor, err := auditing.NewAuditor(logFile)
    if err != nil {
        return nil, err
    }
    return &SecurityManager{auditor: auditor}, nil
}

// HandleSecurity handles input sanitization and logs an audit trail.
func (s *SecurityManager) HandleSecurity(input string, eventDescription string) string {
    // Sanitize input
    sanitizedInput := sanitization.SanitizeInput(input)

    // Record audit event
    s.auditor.RecordAudit(eventDescription)

    // Return sanitized input
    return sanitizedInput
}

// Close closes the auditor associated with the security manager.
func (s *SecurityManager) Close() {
    s.auditor.Close()
}
