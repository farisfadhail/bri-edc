package utils

import (
	"regexp"
	"strings"
)

func CamelToSnake(s string) string {
	re := regexp.MustCompile("(.)([A-Z][a-z]+)")
	s = re.ReplaceAllString(s, "${1}_${2}")
	re = regexp.MustCompile("([a-z0-9])([A-Z])")
	s = re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(s)
}
