package main

import (
	"fmt"
	"strings"
)

func solveDay4(data []string) (int, int) {
	var answer1 int = 0
	var answer2 int = 0
	wordSearch := Map(data, func(item string) []string { return strings.Split(item, "") })

	for row := range len(wordSearch) {
		// fmt.Println(wordSearch[row])
		for col := range len(wordSearch[row]) {
			if wordSearch[row][col] == "X" {
				counter := startSearch(row, col, wordSearch, 1)
				answer1 += counter
			}
		}
	}

	for row := range len(wordSearch) {
		// fmt.Println(wordSearch[row])
		for col := range len(wordSearch[row]) {
			if wordSearch[row][col] == "A" {
				result := checkCross(row, col, wordSearch)
				if result {
					answer2++
				}
			}
		}
	}

	return answer1, answer2
}

const targetWord = "XMAS"

func startSearch(row int, col int, matrix [][]string, targetIndex int) int {
	counter := 0
	directions := [][]int{
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}

	for _, d := range directions {
		newRow := row + d[0]
		newCol := col + d[1]
		fmt.Println("Direction", d)
		if newRow > len(matrix)-1 || newRow < 0 {
			continue
		} else if newCol > len(matrix[newRow])-1 || newCol < 0 {
			continue
		}
		if matrix[newRow][newCol] == string(targetWord[targetIndex]) {
			fmt.Println("Found match", newRow, newCol, row, col)
			result := searchInDirection(newRow, newCol, matrix, targetIndex+1, d)
			if result {
				counter++
			}
		}
	}

	return counter
}

func searchInDirection(row int, col int, matrix [][]string, targetIndex int, direction []int) bool {
	if targetIndex == 4 {
		return true
	}

	newRow := row + direction[0]
	newCol := col + direction[1]
	if newRow > len(matrix)-1 || newRow < 0 {
		return false
	} else if newCol > len(matrix[newRow])-1 || newCol < 0 {
		return false
	}
	fmt.Println("Checking", newRow, newCol)
	if matrix[newRow][newCol] == string(targetWord[targetIndex]) {
		return searchInDirection(newRow, newCol, matrix, targetIndex+1, direction)
	} else {
		fmt.Println("Did not match")
		return false
	}
}

func checkCross(row int, col int, matrix [][]string) bool {
	fmt.Println("Checking cross", row, col)
	candidates := [][]int{
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}
	for _, cand := range candidates {
		newRow := row + cand[0]
		newCol := col + cand[1]
		if newRow > len(matrix)-1 || newRow < 0 {
			return false
		} else if newCol > len(matrix[newRow])-1 || newCol < 0 {
			return false
		}
	}
	configuration := []string{
		matrix[row+candidates[0][0]][col+candidates[0][1]],
		matrix[row+candidates[1][0]][col+candidates[1][1]],
		matrix[row+candidates[2][0]][col+candidates[2][1]],
		matrix[row+candidates[3][0]][col+candidates[3][1]],
	}

	correctConfigs := [][]string{
		{"M", "M", "S", "S"},
		{"S", "S", "M", "M"},
		{"S", "M", "M", "S"},
		{"M", "S", "S", "M"},
	}

	match := true
	for _, cc := range correctConfigs {
		match = true
		for i, _ := range configuration {
			if cc[i] != configuration[i] {
				match = false
				break
			}
		}
		if match {
			fmt.Println("Matched Cross", cc)
			return true
		}
	}
	return false
}
