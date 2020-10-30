package lang

import (
	"crypto/rand"
	"encoding/base64"
)

// StringRandom random string with size given
func StringRandom(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)[0:size]
}
