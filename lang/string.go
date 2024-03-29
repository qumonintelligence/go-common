package lang

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

	case []byte:
		return string(value)

	case primitive.ObjectID:
		return value.Hex()

	default:
		s := fmt.Sprintf("%v", unknown)
		return s
	}
}

// StringLastOf try to get last $count of characters from an unknown
func StringLastOf(unknown interface{}, count int) string {
	if count <= 0 {
		return ""
	}

	_s := StringOf(unknown)
	_l := len(_s) - count
	if _l <= 0 {
		return _s
	}

	return _s[_l:]
}
