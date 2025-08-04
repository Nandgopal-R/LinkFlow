package util

import (
	"strings"
)

func SplitString(str string) []string {
	parts := strings.Split(str, ",")
	return parts
}
