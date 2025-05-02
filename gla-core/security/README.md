# Security Module

This module integrates input sanitization and audit logging, ensuring safe handling of data and keeping track of system interactions. The goal is to prevent security vulnerabilities and provide an audit trail for every significant system event.

## Components

1. **Sanitization**: Functions for sanitizing user inputs to protect against attacks such as SQL injection and cross-site scripting (XSS).
2. **Auditing**: Logging system events and creating an audit trail with timestamps and event descriptions.
3. **SecurityManager**: A wrapper that combines sanitization and auditing functionality, enabling security measures to be applied to inputs while logging events.

## Example Usage

```go
// Initialize the SecurityManager with an audit log file
securityManager, err := security.NewSecurityManager("audit.log")
if err != nil {
    log.Fatal(err)
}
defer securityManager.Close()

// Handle security for a user input and log the event
input := "<script>alert('XSS')</script> DROP TABLE users;"
eventDescription := "User input processed"
sanitizedInput := securityManager.HandleSecurity(input, eventDescription)

fmt.Println(sanitizedInput)
