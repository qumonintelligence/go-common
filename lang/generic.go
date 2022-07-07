package lang

func IsSliceContain[V comparable](slice []V, value V) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}

	return false
}

func ConvertSliceToMap[K comparable, V any](slice []K, value V) map[K]V {
	result := map[K]V{}
	for _, s := range slice {
		result[s] = value
	}

	return result
}
