package util

import (
	"regexp"
	"strings"
)

// ToSlug converts a string to a URL-friendly slug.
func ToSlug(s string) string {
	// Convert the string to lowercase
	s = strings.ToLower(s)

	// Replace all non-alphanumeric characters with hyphens
	reg := regexp.MustCompile("[^a-z0-9]+")
	s = reg.ReplaceAllString(s, "-")

	// Trim hyphens from the start and end of the string
	s = strings.Trim(s, "-")

	// Remove any consecutive hyphens
	s = regexp.MustCompile("-+").ReplaceAllString(s, "-")

	return s
}
