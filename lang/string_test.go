package lang

import "testing"

func TestIsEmptyString(t *testing.T) {

	var s *string

	if !IsEmptyString(s) {
		t.Errorf("IsEmpty failed")
	}

	x := ""
	s = &x

	if !IsEmptyString(s) {
		t.Errorf("IsEmpty failed")
	}
}
