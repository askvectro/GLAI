# Auditing and Logging

This module provides logging and auditing functionality to track and record system activities. It can be used to generate an audit trail of all important actions performed by the system or users.

## Functions

### `RecordAudit`
Logs system events along with timestamps to create an audit trail. Each audit record is stored with the event's description.

### Example Usage
```go
auditor, err := auditing.NewAuditor("audit.log")
if err != nil {
    log.Fatal(err)
}
defer auditor.Close()

auditor.RecordAudit("User logged in.")
