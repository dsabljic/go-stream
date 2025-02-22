package stream

func Map[T, U any](array []T, fn func(T, int, []T) U) []U {
	result := make([]U, len(array))
	for i, v := range array {
		result[i] = fn(v, i, array)
	}

	return result
}
