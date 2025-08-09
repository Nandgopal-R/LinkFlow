package util

import (
	"strings"
)

func SplitString(str string) (string, string, string) {
	parts := strings.Split(str, "_")
	if len(parts) == 2 {
		parts = append(parts, "")
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), strings.TrimSpace(parts[2])
}
