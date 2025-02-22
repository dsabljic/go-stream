package stream

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("Should work with slices of different types", func(t *testing.T) {
		intSlice := []int{1, 2, 3, 4, 5}
		intResult := Map(intSlice, func(v int, i int, arr []int) int {
			return v * 2
		})
		expectedIntResult := []int{2, 4, 6, 8, 10}
		if !reflect.DeepEqual(intResult, expectedIntResult) {
			t.Errorf("Expected %v, but got %v", expectedIntResult, intResult)
		}

		strSlice := []string{"a", "b", "c"}
		strResult := Map(strSlice, func(v string, i int, arr []string) string {
			return v + v
		})
		expectedStrResult := []string{"aa", "bb", "cc"}
		if !reflect.DeepEqual(strResult, expectedStrResult) {
			t.Errorf("Expected %v, but got %v", expectedStrResult, strResult)
		}

		type testStruct struct {
			value int
		}
		structSlice := []testStruct{{1}, {2}, {3}}
		structResult := Map(structSlice, func(v testStruct, i int, arr []testStruct) testStruct {
			return testStruct{v.value + 1}
		})
		expectedStructResult := []testStruct{{2}, {3}, {4}}
		if !reflect.DeepEqual(structResult, expectedStructResult) {
			t.Errorf("Expected %v, but got %v", expectedStructResult, structResult)
		}
	})
}

func TestMapEmptySlice(t *testing.T) {
	input := []int{}
	result := Map(input, func(v, i int, arr []int) int {
		return v * 2
	})

	if len(result) != 0 {
		t.Errorf("Expected empty slice, got slice with length %d", len(result))
	}
}

func TestMapDoesNotModifyOriginalArray(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	Map(original, func(v int, i int, arr []int) int {
		return v * 2
	})

	if !reflect.DeepEqual(original, originalCopy) {
		t.Errorf("Map modified the original array. Expected %v, but got %v", originalCopy, original)
	}
}
