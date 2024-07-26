package common


func SliceContains[T comparable](s []T, val T) int {
	count := 0
	for _, v := range s {
		if v == val {
			count++
		}
	}

	return count
}