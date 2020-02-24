package lang

import (
	"fmt"
)

// IsEmptyString true if given string is nil or empty
func IsEmptyString(s *string) bool {
	return s == nil || len(*s) == 0
}

// StringOf try to get a string of unknown type
func StringOf(unknown interface{}) string {

	if unknown == nil {
		return ""
	}

	switch value := unknown.(type) {
	case string:
		return value

	case *string:
		if value == nil {
			return ""
		}

		return *value

	default:
		s := fmt.Sprintf("%v", unknown)
		return s
	}
}
