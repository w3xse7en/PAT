package problem0028

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
