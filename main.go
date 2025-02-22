package stream

import (
	"fmt"
	"strconv"
)

func main() {

	a := []int{1, 2, 3, 4, 5}

	fmt.Println("Doubled a: ", Map(a, func(el, index int, array []int) int {
		x := array[index] * 2
		return x * el
	}))

	fmt.Println("Int array to string array: ", Map(a, func(el, index int, array []int) string {
		return strconv.Itoa(el)
	}))

	fmt.Println("Even numbers from the array: ", Filter(a, func(i int) bool { return i%2 == 0 }))

	fmt.Println("Sum of all elements of the array: ", Reduce(a, func(acc, i int) int {
		acc += i
		return acc
	}, 0))

	add2 := func(x int) int { return x + 2 }
	multWith5 := func(x int) int { return x * 5 }
	sub3 := func(x int) int { return x - 3 }
	sumTwoNumbers := func(a, b int) int { return a + b }
	addAndMultiply := func(a, b int) (int, int) { return a + b, a * b }
	parseStringToIntDoubleReturnVals := func(s string) (int, int) { i, _ := strconv.Atoi(s); return i, 10 } // toy example
	// parseStringToInt := func(s string) (int, error) { return strconv.Atoi(s) }

	pipe1 := Pipe(sub3, multWith5, add2)
	result := pipe1(5)[0]
	fmt.Println("add2(multWith5(sub3(5) = 2) = 10) = ", result)

	pipe2 := Pipe(parseStringToIntDoubleReturnVals, addAndMultiply, sumTwoNumbers)
	result = pipe2("5")[0]
	fmt.Println("sumTwoNumbers(addAndMultiply(parseStringToIntDoubleReturnVals(\"5\") = 5, 10) = 15, 50) = ", result)
}
