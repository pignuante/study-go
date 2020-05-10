package utils

import "strings"

// CleanString Trim Space
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
