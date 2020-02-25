package lang

import (
	"testing"

	"gotest.tools/assert"
)

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

func TestStringLastOfEmpty(t *testing.T) {
	s := ""

	assert.Equal(t, StringLastOf(s, 3), "", "Should be empty")
}

func TestStringLastOfNegative(t *testing.T) {
	s := "abc"
	assert.Equal(t, StringLastOf(s, -1), "", "Shoud be empty")
}

func TestStringLastOfAll(t *testing.T) {
	s := "abc"
	assert.Equal(t, StringLastOf(s, 100), "abc", "Should return all")
}

func TestStringLast2(t *testing.T) {
	s := "abc"
	assert.Equal(t, StringLastOf(s, 2), "bc", "Should return last 2 chars")
}
