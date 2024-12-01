package main

import (
	"fmt"
	"os"
	"strconv"
)

func checkArgs(args []string) {
	if len(args) < 1 {
		fmt.Println("Missing 'day' argument!")
		os.Exit(-1)
	} else if len(args) > 1 {
		fmt.Println("Too many arguments!")
		os.Exit(-1)
	}
}

func main() {
	solvers := []func([]string) int{
		solveDay1,
	}

	args := os.Args[1:]
	checkArgs(args)
	day, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Running Day %v Function...\n", day)

	inputFile := fmt.Sprintf("day%v.txt", 1)
	result := readFileByDelimiter(inputFile, "\n")

	solvers[day-1](result)
}
