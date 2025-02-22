package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestPipePanicOnNoFunctions(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Pipe did not panic when no functions were provided")
		} else if r != "Pipe: no functions provided" {
			t.Errorf("Unexpected panic message: %v", r)
		}
	}()

	Pipe()
}

func TestPipePanicOnNonFunction(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Pipe did not panic when a non-function argument was provided")
		} else if r != "Pipe: argument 0 is not a function" {
			t.Errorf("Unexpected panic message: %v", r)
		}
	}()

	Pipe("not a function")
}

func TestPipePanicOnMismatchedFunctions(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Pipe did not panic when functions had mismatched inputs/outputs")
		} else if msg, ok := r.(string); !ok || !strings.Contains(msg, "Pipe: function 0 returns 1 values, but function 1 expects 2 parameters") {
			t.Errorf("Unexpected panic message: %v", r)
		}
	}()

	fn1 := func(x int) int { return x * 2 }
	fn2 := func(x, y int) int { return x + y }
	Pipe(fn1, fn2)
}

func TestPipePanicOnTypeMismatch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Pipe did not panic when there was a type mismatch")
		} else if r != "Pipe: type mismatch between function 0 output #0 (int) and function 1 input #0 (string)" {
			t.Errorf("Unexpected panic message: %v", r)
		}
	}()

	f1 := func(x int) int { return x + 1 }
	f2 := func(s string) string { return s + "!" }

	Pipe(f1, f2)
}

func TestPipePanicOnArgumentMismatch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Pipe did not panic when argument count mismatched")
		} else if r != "Pipe: first function expects 2 arguments, got 1" {
			t.Errorf("Unexpected panic message: %v", r)
		}
	}()

	fn1 := func(a, b int) int { return a + b }
	pipedFunc := Pipe(fn1)
	pipedFunc(1) // This should panic as fn1 expects 2 arguments
}

func TestPipeCorrectlyChainsFunctions(t *testing.T) {
	// Define test functions
	addOne := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	toString := func(x int) string { return fmt.Sprintf("%d", x) }

	// Create a pipeline of functions
	pipeline := Pipe(addOne, double, toString)

	// Execute the pipeline with an initial input
	result := pipeline(5)

	// Check if the result is correct
	expected := []interface{}{"12"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestPipepleInputOutput(t *testing.T) {
	add := func(a, b int) int { return a + b }
	multiply := func(x int) (int, int) { return x * 3, x * 2 }
	subtract := func(a, b int) int { return a - b }

	pipeline := Pipe(add, multiply, subtract)

	result := pipeline(3, 7)

	expected := []interface{}{10}
	if len(result) != len(expected) {
		t.Errorf("Expected %d results, but got %d", len(expected), len(result))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestPipeWithNoInputParameters(t *testing.T) {
	getOne := func() int { return 1 }
	addTwo := func(x int) int { return x + 2 }
	toString := func(x int) string { return fmt.Sprintf("%d", x) }

	pipeline := Pipe(getOne, addTwo, toString)
	result := pipeline()

	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result))
	}

	if str, ok := result[0].(string); !ok || str != "3" {
		t.Errorf("Expected result to be string '3', got %v", result[0])
	}
}

func TestPipeWithNoOutputParameters(t *testing.T) {
	callCount := 0
	noOutputFunc := func() {
		callCount++
	}

	piped := Pipe(noOutputFunc)
	result := piped()

	if len(result) != 0 {
		t.Errorf("Expected empty result slice, got %v", result)
	}

	if callCount != 1 {
		t.Errorf("Expected function to be called once, but it was called %d times", callCount)
	}
}

func TestPipeCorrectOutput(t *testing.T) {
	// Define test functions
	double := func(x int) int { return x * 2 }
	addOne := func(x int) int { return x + 1 }
	toString := func(x int) string { return fmt.Sprintf("%d", x) }

	// Create a pipe with the test functions
	pipe := Pipe(double, addOne, toString)

	// Execute the pipe with an initial value
	result := pipe(5)

	// Check if the result is correct
	expected := []interface{}{"11"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
