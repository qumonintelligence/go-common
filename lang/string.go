package lang

// IsEmptyString true if given string is nil or empty
func IsEmptyString(s *string) bool {
	return s == nil || len(*s) == 0
}
