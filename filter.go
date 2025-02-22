package main

func Filter[T any](array []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range array {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}
