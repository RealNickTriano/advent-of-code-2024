package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

func checkConditional(file *os.File) string {
	doVerb := "o()"
	dontVerb := "on't()"
	stringPattern := "o"
	currentIndex := 0

	verb := doVerb

	b1 := make([]byte, 1)
	for {
		_, err := file.Read(b1)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println("reading for:", string(b1[0]))
		if currentIndex > 1 {
			if b1[0] == verb[currentIndex] {
				fmt.Println("Matched in verb", string(b1[0]), string(verb[currentIndex]), currentIndex)
				currentIndex++
				if currentIndex == len(verb) {
					if len(verb) == 3 {
						return "do"
					} else {
						return "dont"
					}
				}
			} else {
				return ""
			}
		} else if currentIndex == 1 {
			if b1[0] == '(' {
				fmt.Println("Verb is do")
				verb = doVerb
				currentIndex++
			} else if b1[0] == 'n' {
				fmt.Println("Verb is dont")
				verb = dontVerb
				currentIndex++
			} else {
				return ""
			}
		} else if b1[0] == stringPattern[currentIndex] {
			fmt.Println("Matched", string(b1[0]))
			currentIndex++
		} else {
			return ""
		}

	}
	return ""
}

func solveDay3(data []string) (int, int) {
	var answer1 int = 0
	var answer2 int = 0

	file, err := os.Open("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b1 := make([]byte, 1)
	stringPattern := "mul(x,x)"
	currentIndex := 0
	num1 := ""
	num2 := ""
	lookingForNums := false
	commaFound := false
	allowMul := true

	resetFunc := func() {
		lookingForNums = false
		commaFound = false
		currentIndex = 0
		num1 = ""
		num2 = ""
	}

	for {
		_, err := file.Read(b1)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if b1[0] == 'd' {
			// Possible conditional here
			fmt.Println("Found d")
			result := checkConditional(file)
			if result == "do" {
				fmt.Println("allow mul")
				allowMul = true
			} else if result == "dont" {
				fmt.Println("dont allow mul")
				allowMul = false
			}
			resetFunc()
			continue
		}

		fmt.Println("Looking for:", string(stringPattern[currentIndex]))

		if stringPattern[currentIndex] == 'x' {
			lookingForNums = true
		}

		if lookingForNums {
			if unicode.IsDigit(rune(b1[0])) {
				fmt.Println("Found digit:", string(b1[0]))
				if !commaFound {
					fmt.Println("Adding to num1", string(b1[0]))
					num1 += string(b1[0])
				} else {
					fmt.Println("Adding to num2", string(b1[0]))
					num2 += string(b1[0])
				}
			} else if b1[0] == ',' {
				fmt.Println("Found comma:", string(b1[0]))
				commaFound = true
			} else if b1[0] == ')' {
				if allowMul {
					fmt.Println("Multiplying...", num1, num2)
					answer1 += stringToInt(num1) * stringToInt(num2)
				}
				resetFunc()
			} else {
				resetFunc()
			}
		} else if b1[0] == stringPattern[currentIndex] {
			currentIndex++
			fmt.Println(string(b1[0]))

		} else {
			resetFunc()
		}
	}

	return answer1, answer2
}
