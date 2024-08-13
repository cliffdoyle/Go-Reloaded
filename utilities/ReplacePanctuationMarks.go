package utilities


import (
	"regexp"
	"strings"
)

func FormatText(text string) string {
	// Define regular expression for punctuations
	punctuationRegex := regexp.MustCompile(`\s+([.,!?;:]+)`)
	pattern2 := regexp.MustCompile(`([,.;:?!]+)(\w+)`)

	// Replace punctuations with spaces around them
	text = punctuationRegex.ReplaceAllString(text, "$1")
	text2 := pattern2.ReplaceAllString(text, "$1 $2")

	return text2
}

func Quotation(s string) string {
	s = strings.ReplaceAll(s, "' ", " '")
	s = strings.ReplaceAll(s, " '", "'")

	return strings.TrimSpace(s)
}
