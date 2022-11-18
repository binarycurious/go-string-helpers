package stringhelpers

import "regexp"

// IsEmptyOrWhitespace : a simple function to test if a string is empty or only whitespace chars
func IsEmptyOrWhitespace(s *string) bool {
	if s == nil {
		return true
	}
	emptyMatch, _ := regexp.MatchString("^\\s*$", string(*s))
	return emptyMatch
}
