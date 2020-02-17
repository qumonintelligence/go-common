package lang

import (
	"encoding/json"
	"fmt"
	"strings"
)

// IMapDecoder map decoder interface
type IMapDecoder interface {
	Decode(in interface{}, out interface{}) error
}

// JSONMapDecoder decode map to an object
type JSONMapDecoder struct {
}

// Decode in into out
func (d *JSONMapDecoder) Decode(in interface{}, out interface{}) error {
	js, _ := json.Marshal(in)
	return json.Unmarshal(js, out)
}

var defaultDecoder IMapDecoder

func init() {
	defaultDecoder = &JSONMapDecoder{}
}

// StringMap with additional functions
type StringMap map[string]interface{}

// Get value by key
func (m *StringMap) Get(key string) (interface{}, bool) {
	val, ok := (*m)[key]
	return val, ok
}

// Set key = value
func (m *StringMap) Set(key string, value interface{}) {
	(*m)[key] = value
}

// GetAs decode using json
func (m *StringMap) GetAs(key string, out interface{}) error {
	return m.GetAsWith(defaultDecoder, key, out)
}

// GetAsWith given decoder
func (m *StringMap) GetAsWith(decoder IMapDecoder, key string, out interface{}) error {
	if val, ok := m.Get(key); ok {
		return decoder.Decode(val, out)
	}

	return fmt.Errorf("Not found: %v", key)
}

// GetAsByPath Get a value with dot-separated path
func (m *StringMap) GetAsByPath(path string, out interface{}) error {

	// split it by .
	paths := strings.Split(path, ".")

	var x interface{}
	x = *m

	for _, p := range paths {

		if i, ok := x.(StringMap); ok {
			if val, ok := i.Get(p); ok {
				x = val
				continue
			}

			return fmt.Errorf("Not found: %v", p)
		}

		if i, ok := x.(map[string]interface{}); ok {
			if val, ok := i[p]; ok {
				x = val
				continue
			}

			return fmt.Errorf("Not found: %v", p)
		}

		// unknown type, failed to resolve next
		return fmt.Errorf("Unknown type: %T", x)
	}

	if x != nil {
		return defaultDecoder.Decode(x, out)
	}

	return nil
}
