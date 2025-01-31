package twlib

import (
	"fmt"
	"mime/multipart"
	"strings"
)

// TrimQuotes trims double, single, and backtick quotes from a string.
func TrimQuotes(s string) string {
	if len(s) >= 2 {
		lastChar := s[len(s)-1]
		if lastChar == s[0] && (lastChar == '"' || lastChar == '\'' || lastChar == '`') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// GetStructuredFilename extracts the filename from a multipart file header.
// It looks for the "filename" field in the "Content-Disposition" header and
// returns the value, which should include the directory structure if provided.
func GetStructuredFilename(file *multipart.FileHeader) (string, error) {
	for _, content := range strings.Split(file.Header.Get("Content-Disposition"), ";") {
		parts := strings.Split(strings.TrimSpace(content), "=")
		if len(parts) != 2 {
			continue
		}
		if parts[0] == "filename" {
			return TrimQuotes(parts[1]), nil
		}
	}
	return "", fmt.Errorf("filename not found")
}
