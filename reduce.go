package main

func Reduce[T any](array []T, fn func(T, T) T, initVal T) T {
	result := initVal

	for _, v := range array {
		result = fn(result, v)
	}

	return result
}
