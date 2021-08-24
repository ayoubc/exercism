// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	r, _ := regexp.Compile(`[\s,_-]+`)
	s = r.ReplaceAllString(s, " ")
	parts := strings.Split(s, " ")
	res := ""
	for _, part := range parts {
		res += strings.ToUpper(string(part[0]))
	}
	return res
}
