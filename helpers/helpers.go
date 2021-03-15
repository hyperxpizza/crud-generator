package helpers

import "strings"

func Between(value, a, b string) string {
	value = strings.ToLower(value)
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}

	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}

	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func CheckIfContains(s string, list []string) string {
	for _, v := range list {
		if strings.Contains(s, v) {
			return v
		}
	}

	return ""
}
