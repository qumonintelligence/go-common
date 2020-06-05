package lang

import "time"

// ToDurationPtr convert a time.Duration to a pointer
func ToDurationPtr(d time.Duration) *time.Duration {
	return &d
}
