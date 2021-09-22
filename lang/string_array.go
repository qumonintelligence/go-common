package lang

import "fmt"

// StringArrayOf try to convert an unknown to string array
func StringArrayOf(unknown interface{}) ([]string, error) {
	if unknown == nil {
		return []string{}, fmt.Errorf("unknown is nil")
	}

	switch value := unknown.(type) {
	case []interface{}:
		values := []string{}
		for _, k := range value {
			values = append(values, StringOf(k))
		}
		return values, nil

	case []string:
		return value, nil

	default:
		return []string{}, fmt.Errorf("unsupported type %T", unknown)
	}
}

func StringSliceContains(slices []string, value string) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}

	return false
}
