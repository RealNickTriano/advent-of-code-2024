package main

import "fmt"

func main() {
	println("Hello World")
	result := readFileByDelimiter("input_1.txt", "\n")
	fmt.Println(result)
}
