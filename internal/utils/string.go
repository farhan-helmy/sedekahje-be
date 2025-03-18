package utils

import (
	"regexp"
	"strings"
)

func Slugify(str string) string {
	// Convert to lowercase
	str = strings.ToLower(str)

	// Replace spaces with hyphens
	str = strings.ReplaceAll(str, " ", "-")

	// Remove all non-alphanumeric characters except hyphens
	reg := regexp.MustCompile("[^a-z0-9-]")
	str = reg.ReplaceAllString(str, "")

	return str
}
