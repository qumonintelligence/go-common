package lang

import (
	"fmt"
	"strconv"
)

// Int64Of try to get an int64 of unknown type
func Int64Of(unknown interface{}) (int64, error) {

	if unknown == nil {
		return 0, fmt.Errorf("Value is nil")
	}

	switch value := unknown.(type) {
	case int64:
		return value, nil

	case int:
		return int64(value), nil

	case int32:
		return int64(value), nil

	case string:
		return strconv.ParseInt(value, 10, 64)

	default:
		s := fmt.Sprintf("%v", unknown)
		return strconv.ParseInt(s, 10, 64)
	}
}
