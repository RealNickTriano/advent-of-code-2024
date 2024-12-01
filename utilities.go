package main

import (
	"fmt"
	"strconv"
)

type Slice[T any] []T

func PrintSlice[T any](s Slice[T]) {
	fmt.Print("[")
	for i, v := range s {
		end := ""
		if i != len(s)-1 {
			end = ", "
		}
		fmt.Printf("%v%s", v, end)
	}
	fmt.Println("]")
}

func PrintSliceWithNewLine[T any](s Slice[T]) {
	for _, v := range s {
		end := "\n"
		fmt.Printf("%v%s", v, end)
	}
}

func stringToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func Map[T, V any](list []T, fn func(T) V) []V {
	result := make([]V, len(list))
	for i, oldVal := range list {
		result[i] = fn(oldVal)
	}
	return result
}

func Filter[T any](list []T, fn func(T) bool) []T {
	result := make([]T, 0, len(list))
	count := 0
	for _, val := range list {
		if fn(val) {
			count++
			result = append(result, val)
		}
	}
	return result
}
