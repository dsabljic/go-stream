package main

import (
	"fmt"
	"reflect"
)

// func Pipe[T any](fns ...func(T) T) func(T) T {
// 	return func(x T) T {
// 		result := x
// 		for _, fn := range fns {
// 			result = fn(result)
// 		}

// 		return result
// 	}
// }

func Pipe(fns ...interface{}) func(args ...interface{}) []interface{} {
	if len(fns) == 0 {
		panic("Pipe: no functions provided")
	}

	fnTypes := make([]reflect.Type, len(fns))
	for i, fn := range fns {
		t := reflect.TypeOf(fn)
		if t.Kind() != reflect.Func {
			panic(fmt.Sprintf("Pipe: argument %d is not a function", i))
		}
		fnTypes[i] = t

		if i > 0 {
			prevOutCount := fnTypes[i-1].NumOut()
			currentInCount := t.NumIn()
			if prevOutCount != currentInCount {
				panic(fmt.Sprintf("Pipe: function %d returns %d values, but function %d expects %d parameters",
					i-1, prevOutCount, i, currentInCount))
			}
			for j := 0; j < currentInCount; j++ {
				prevOutType := fnTypes[i-1].Out(j)
				currInType := t.In(j)
				if !prevOutType.AssignableTo(currInType) {
					panic(fmt.Sprintf("Pipe: type mismatch between function %d output #%d (%v) and function %d input #%d (%v)",
						i-1, j, prevOutType, i, j, currInType))
				}
			}
		}
	}

	return func(args ...interface{}) []interface{} {
		firstFnType := fnTypes[0]
		if len(args) != firstFnType.NumIn() {
			panic(fmt.Sprintf("Pipe: first function expects %d arguments, got %d", firstFnType.NumIn(), len(args)))
		}

		currentVals := make([]reflect.Value, len(args))
		for i, arg := range args {
			currentVals[i] = reflect.ValueOf(arg)
		}

		for i, fn := range fns {
			fnVal := reflect.ValueOf(fn)
			if fnVal.Type().NumIn() != len(currentVals) {
				panic(fmt.Sprintf("Pipe: function %d expected %d arguments but got %d", i, fnVal.Type().NumIn(), len(currentVals)))
			} // might be redundant

			results := fnVal.Call(currentVals)
			currentVals = results
		}

		out := make([]interface{}, len(currentVals))
		for i, v := range currentVals {
			out[i] = v.Interface()
		}

		return out
	}
}
