package lang

import (
	"fmt"
	"strconv"
)

// IntOf try to get an int64 of unknown type
func IntOf(unknown interface{}) (int, error) {

	if unknown == nil {
		return 0, fmt.Errorf("Value is nil")
	}

	switch value := unknown.(type) {
	case int64:
		return int(value), nil

	case int:
		return value, nil

	case int32:
		return int(value), nil

	case string:
		return strconv.Atoi(value)

	case []byte:
		s := string(value)
		return strconv.Atoi(s)

	default:
		s := fmt.Sprintf("%v", unknown)
		return strconv.Atoi(s)
	}
}
