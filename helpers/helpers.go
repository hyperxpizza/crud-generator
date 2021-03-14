package helpers

import (
	"regexp"

	"log"
)

func RemoveEmptyLines(s string) (string, error) {
	r, err := regexp.Compile("\n\n")
	if err != nil {
		log.Fatalf("regexp.Compile failed: %v\n", err)
		return "", err
	}

	s1 := r.ReplaceAllString(s, "")
	return s1, nil
}
