package lang

import "hash/fnv"

// HashOf calculate uint32 hash code of a string
func HashOf(value string) uint32 {
	hasher := fnv.New32()
	hasher.Write([]byte(value))
	return hasher.Sum32()
}
