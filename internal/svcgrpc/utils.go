package svcgrpc

import (
	"regexp"
	"strings"
)

var entityRegExp *regexp.Regexp
var patternRegExp *regexp.Regexp

func validEntity(name string) bool {
	return entityRegExp.MatchString(name)
}

func validScope(name string) bool {
	return entityRegExp.MatchString(name)
}

func validPattern(pattern string) bool {
	return patternRegExp.MatchString(pattern)
}

func EscapePattern(pattern string) string {
	pattern = strings.ReplaceAll(pattern, "_", "\\_")
	pattern = strings.ReplaceAll(pattern, "%", "\\%")
	pattern = strings.ReplaceAll(pattern, "*", "%")
	return pattern
}

func init() {
	entityRegExp = regexp.MustCompile("^(?:[A-Za-z0-9_-]+(?:[:/][A-Za-z0-9_-]+)*)$")
	patternRegExp = regexp.MustCompile(`^(?:[A-Za-z0-9_-]+(?:[:/]([A-Za-z0-9_-]+|\*))*)$`)
}
