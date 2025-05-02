# Input Sanitization and Validation

This module provides functions to sanitize and validate user inputs to prevent injection attacks and ensure safe handling of data within the system.

## Functions

### `SanitizeInput`
Cleanses input to remove potentially harmful content, including:

- Script tags
- SQL injection patterns

This helps mitigate cross-site scripting (XSS) and SQL injection vulnerabilities.

## Example Usage
```go
input := "<script>alert('XSS')</script> OR 1=1"
sanitized := sanitization.SanitizeInput(input)
fmt.Println(sanitized) // Output: alert('XSS') OR 1=1
