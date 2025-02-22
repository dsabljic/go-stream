package main

import (
	"testing"
)

func TestReduceWithIntegerAddition(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Reduce(numbers, func(a, b int) int { return a + b }, 0)
	expected := 15

	if sum != expected {
		t.Errorf("Reduce with integer addition: expected %d, got %d", expected, sum)
	}
}

func TestReduceEmptyArray(t *testing.T) {
	emptyArray := []int{}
	initialValue := 42
	result := Reduce(emptyArray, func(a, b int) int { return a + b }, initialValue)
	if result != initialValue {
		t.Errorf("Expected %d, but got %d", initialValue, result)
	}
}

func TestReduceString(t *testing.T) {
	input := []string{"Hello", " ", "World", "!"}
	expected := "Hello World!"

	result := Reduce(input, func(a, b string) string {
		return a + b
	}, "")

	if result != expected {
		t.Errorf("Reduce failed. Expected %s, but got %s", expected, result)
	}
}

func TestReduceWithSingleElement(t *testing.T) {
	array := []int{5}
	initVal := 0
	result := Reduce(array, func(a, b int) int { return a + b }, initVal)
	expected := 5

	if result != expected {
		t.Errorf("Reduce() with single element: got %v, want %v", result, expected)
	}
}

func TestReduceWithCustomStructAndFunction(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	oldestPerson := Reduce(people, func(a, b Person) Person {
		if a.Age > b.Age {
			return a
		}
		return b
	}, Person{"", 0})

	expected := Person{"Charlie", 35}
	if oldestPerson != expected {
		t.Errorf("Expected oldest person to be %v, but got %v", expected, oldestPerson)
	}
}

func TestReduceWithLargeArray(t *testing.T) {
	// Create a large array
	size := 1000000
	largeArray := make([]int, size)
	for i := 0; i < size; i++ {
		largeArray[i] = i
	}

	// Define a simple sum function
	sum := func(a, b int) int {
		return a + b
	}

	// Perform reduction
	result := Reduce(largeArray, sum, 0)

	// Calculate expected sum
	expected := size * (size - 1) / 2

	// Check if the result matches the expected sum
	if result != expected {
		t.Errorf("Reduce failed for large array. Expected %d, got %d", expected, result)
	}
}

func TestReduceWithNegativeNumbers(t *testing.T) {
	numbers := []int{-5, -3, -8, -2, -1}
	expected := -19

	result := Reduce(numbers, func(a, b int) int {
		return a + b
	}, 0)

	if result != expected {
		t.Errorf("Reduce with negative numbers: expected %d, but got %d", expected, result)
	}
}

func TestReduceNilFunction(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	var nilFunc func(int, int) int

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Reduce to panic with nil function, but it didn't")
		}
	}()

	Reduce(array, nilFunc, 0)
}

func TestReduceNonCommutative(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	subtract := func(a, b int) int {
		return a - b
	}

	result := Reduce(numbers, subtract, 100)

	expected := 90 // 100 - 1 - 2 - 3 - 4
	if result != expected {
		t.Errorf("Reduce failed to preserve order of operations. Expected %d, got %d", expected, result)
	}
}
