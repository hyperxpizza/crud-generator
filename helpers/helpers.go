package helpers

import (
	"regexp"

	"log"
)

func RemoveEmptyLine(s string) (string, error) {
	r, err := regexp.Compile("\n\n")
	if err != nil {
		log.Fatalf("regexp.Compile failed: %v\n", err)
		return "", err
	}

	s1 := r.ReplaceAllString(s, "\n")
	return s1, nil
}
