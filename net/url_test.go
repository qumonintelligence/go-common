package net

import (
	"testing"
	"gotest.tools/assert"
)

func TestURL (t *testing.T){
	s, err := RequestURI("https://www.google.com/abcedef?xyz")
	assert.NilError(t,err)
	t.Logf("Test output: %v",s)
}
