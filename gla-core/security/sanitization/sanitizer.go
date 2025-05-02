package sanitization

import (
	"strings"
	"regexp"
)

// SanitizeInput removes dangerous characters from input to prevent injections or malicious code.
func SanitizeInput(input string) string {
	// Example: strip out script tags and SQL injection patterns
	input = removeScriptTags(input)
	input = removeSQLInjectionPatterns(input)
	return input
}

// removeScriptTags removes potential script tags from input
func removeScriptTags(input string) string {
	re := regexp.MustCompile(`(?i)<script.*?>.*?</script>`)
	return re.ReplaceAllString(input, "")
}

// removeSQLInjectionPatterns filters out common SQL injection patterns
func removeSQLInjectionPatterns(input string) string {
	re := regexp.MustCompile(`(?i)(union.*select|--|;|#)`)
	return re.ReplaceAllString(input, "")
}
