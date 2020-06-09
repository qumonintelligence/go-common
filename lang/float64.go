package lang

import (
	"fmt"
	"strconv"
)

// Float64Of try to get an float64 of unknown type
func Float64Of(unknown interface{}) (float64, error) {

	if unknown == nil {
		return 0, fmt.Errorf("Value is nil")
	}

	switch value := unknown.(type) {
	case int64:
		return float64(value), nil

	case int:
		return float64(value), nil

	case int32:
		return float64(value), nil

	case string:
		return strconv.ParseFloat(value, 64)

	case float32:
		return float64(value), nil

	case float64:
		return float64(value), nil

	case []byte:
		s := string(value)
		return strconv.ParseFloat(s, 64)

	default:
		s := fmt.Sprintf("%v", unknown)
		return strconv.ParseFloat(s, 64)
	}
}
