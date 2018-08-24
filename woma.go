package woma

import (
	"strings"
)

// Normalize converts the given string to lower case without leading and
// trailing whitespaces.
func Normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
