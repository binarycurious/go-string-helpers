package stringhelpers

// Coalesces : return the first string that is not empty or an empty string if all items are empty
func Coalesces(ss ...string) *string {
	for _, s := range ss {
		if s == "" {
			continue
		}
		return &s
	}
	e := ""
	return &e
}

// CoalesceWhitespace : return the first string containing non whitespace chars or an empty string if no match
func CoalesceWhitespace(ss ...string) *string {
	for _, s := range ss {
		if IsEmptyOrWhitespace(&s) {
			continue
		}
		return &s
	}
	e := ""
	return &e
}
