package stream

import (
	"testing"
)

func TestEmptySliceFilter(t *testing.T) {
	input := []int{}
	expected := []int{}
	result := Filter(input, func(i int) bool { return i > 0 })

	if len(result) != len(expected) {
		t.Errorf("expected slice of length %d, got %d", len(expected), len(result))
	}

	if len(result) > 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func TestFilterInts(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := Filter(input, func(i int) bool { return i%2 == 0 })

	if len(result) != len(expected) {
		t.Errorf("expected slice of length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("at index %d: expected %d, got %d", i, expected[i], v)
		}
	}
}

func TestFilterAlwaysFalse(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 0
	result := Filter(input, func(i int) bool { return false })

	if len(result) != expected {
		t.Errorf("expected slice of length %d, got %d", expected, len(result))
	}
}

func TestFilterAllTrue(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	result := Filter(input, func(i int) bool { return true })

	if len(result) != len(expected) {
		t.Errorf("expected slice of length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("at index %d: expected %d, got %d", i, expected[i], v)
		}
	}
}

func TestFilterSingleElement(t *testing.T) {
	input := []int{5}
	expected := []int{5}
	result := Filter(input, func(i int) bool { return i > 0 })

	if len(result) != len(expected) {
		t.Errorf("expected slice of length %d, got %d", len(expected), len(result))
	}

	if result[0] != expected[0] {
		t.Errorf("expected %d, got %d", expected[0], result[0])
	}

	// Test with a false condition
	resultFalse := Filter(input, func(i int) bool { return i < 0 })
	if len(resultFalse) != 0 {
		t.Errorf("expected empty slice, got %v", resultFalse)
	}
}

func TestFilterCustomStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	input := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
		{"David", 40},
	}

	expected := []Person{
		{"Charlie", 35},
		{"David", 40},
	}

	result := Filter(input, func(p Person) bool { return p.Age >= 35 })

	if len(result) != len(expected) {
		t.Errorf("expected slice of length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("at index %d: expected %+v, got %+v", i, expected[i], v)
		}
	}
}

func TestFilterMaintainsOrder(t *testing.T) {
	input := []int{1, 4, 2, 5, 3}
	expected := []int{4, 5}
	result := Filter(input, func(i int) bool { return i > 3 })

	if len(result) != len(expected) {
		t.Errorf("expected slice of length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("at index %d: expected %d, got %d", i, expected[i], v)
		}
	}
}

func TestFilterInterfaceSlice(t *testing.T) {
	input := []interface{}{1, "two", 3.14, true, "five"}
	expected := []interface{}{1, 3.14, true}
	result := Filter(input, func(v interface{}) bool {
		_, ok := v.(string)
		return !ok
	})

	if len(result) != len(expected) {
		t.Errorf("expected slice of length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("at index %d: expected %v, got %v", i, expected[i], v)
		}
	}
}

func TestFilterDoesNotModifyInput(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	original := make([]int, len(input))
	copy(original, input)

	Filter(input, func(i int) bool { return i%2 == 0 })

	if len(input) != len(original) {
		t.Errorf("Filter modified the length of the input slice. Expected: %d, Got: %d", len(original), len(input))
	}

	for i, v := range input {
		if v != original[i] {
			t.Errorf("Filter modified the input slice at index %d. Expected: %d, Got: %d", i, original[i], v)
		}
	}
}
