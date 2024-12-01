package main

import "fmt"

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
