package lang

import (
	"fmt"
)

// Uint64Of try to get an uint64 of unknown type
func Uint64Of(unknown interface{}) (uint64, error) {

	// sanity checking
	if unknown == nil {
		return 0, fmt.Errorf("Value is nil")
	}

	switch value := unknown.(type) {
	case uint64:
		return value, nil

	default:
		r, e := Int64Of(unknown)
		return uint64(r), e
	}
}
