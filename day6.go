package main

import (
	"fmt"
	"strings"
)

func solveDay6(data []string) (int, int) {
	answer1 := 0
	answer2 := 0

	directionMap := map[int][]int{
		0:   {-1, 0},
		90:  {0, 1},
		180: {1, 0},
		270: {0, -1},
	}
	degreesTurned := 0

	mappedArea := Map(data, func(item string) []string { return strings.Split(item, "") })

	guardRow, guardCol := findStart(mappedArea)
	// 32: true mappedArea[3][2] visited
	visited := make(map[string]bool)

	for {
		// if off the map end loop
		if outOfBounds(mappedArea, guardRow, guardCol) {
			break
		}

		fmt.Println("now at", guardRow, guardCol)
		posRowInfront := guardRow + directionMap[degreesTurned][0]
		posColInfront := guardCol + directionMap[degreesTurned][1]

		if !outOfBounds(mappedArea, posRowInfront, posColInfront) && mappedArea[posRowInfront][posColInfront] == "#" {
			// Turn
			fmt.Println(posRowInfront, posColInfront, "Blocked")
			degreesTurned += 90
			if degreesTurned == 360 {
				degreesTurned = 0
			}
			fmt.Println("Now facing", degreesTurned)
			continue
		} else {
			// set visited and move
			encode := string(guardRow) + string(guardCol)
			fmt.Println("visited", encode)
			visited[encode] = true
			guardRow = guardRow + directionMap[degreesTurned][0]
			guardCol = guardCol + directionMap[degreesTurned][1]
		}
	}

	answer1 = len(visited)
	answer2 = partTwo(data)
	return answer1, answer2

}

func copyTwoDSlice(sliceToCopy [][]string) [][]string {
	newSlice := make([][]string, len(sliceToCopy))
	copy(newSlice, sliceToCopy)
	for i, _ := range newSlice {
		copy(newSlice[i], sliceToCopy[i])
	}
	return newSlice
}

func partTwo(data []string) int {
	answer2 := 0

	directionMap := map[int][]int{
		0:   {-1, 0},
		90:  {0, 1},
		180: {1, 0},
		270: {0, -1},
	}
	degreesTurned := 0

	mappedArea1 := Map(data, func(item string) []string { return strings.Split(item, "") })
	mappedArea := make([][]string, len(mappedArea1))
	copy(mappedArea, mappedArea1)

	guardRow, guardCol := findStart(mappedArea)
	// 32: true mappedArea[3][2] visited
	visited := make(map[string]bool)
	size := len(mappedArea) * len(mappedArea[0])
	obstacleRow := 0
	obstacleCol := 0
	for x := range size {
		fmt.Println("Test", obstacleRow, obstacleCol, x)
		// reset map
		mappedArea = copyTwoDSlice(mappedArea1)
		PrintSlice(mappedArea)
		if mappedArea[obstacleRow][obstacleCol] == "^" {
			continue
		}
		// test obstacle
		mappedArea[obstacleRow][obstacleCol] = "O"

		for {
			// if off the map end loop
			if outOfBounds(mappedArea, guardRow, guardCol) {
				break
			}

			fmt.Println("now at", guardRow, guardCol)
			posRowInfront := guardRow + directionMap[degreesTurned][0]
			posColInfront := guardCol + directionMap[degreesTurned][1]

			if !outOfBounds(mappedArea, posRowInfront, posColInfront) && (mappedArea[posRowInfront][posColInfront] == "#" || mappedArea[posRowInfront][posColInfront] == "O") {
				if mappedArea[posRowInfront][posColInfront] == "O" {
					answer2++
				}
				// Turn
				fmt.Println(posRowInfront, posColInfront, "Blocked")
				degreesTurned += 90
				if degreesTurned == 360 {
					degreesTurned = 0
				}
				fmt.Println("Now facing", degreesTurned)
				continue
			} else {
				// set visited and move
				encode := string(guardRow) + string(guardCol)
				fmt.Println("visited", encode)
				visited[encode] = true
				guardRow = guardRow + directionMap[degreesTurned][0]
				guardCol = guardCol + directionMap[degreesTurned][1]
			}
		}
		obstacleCol++
		if obstacleCol > len(mappedArea[obstacleRow])-1 {
			obstacleCol = 0
			obstacleRow++
		}
	}

	return answer2

}

func outOfBounds(matrix [][]string, row int, col int) bool {
	if row > len(matrix)-1 || row < 0 {
		return true
	} else if col > len(matrix[row])-1 || col < 0 {
		return true
	}
	return false
}

func findStart(area [][]string) (int, int) {
	for row, _ := range area {
		for col, _ := range area[row] {
			if area[row][col] == "^" {
				return row, col
			}
		}
	}

	panic("No Start")
}

func printGuardMap(area [][]string, pos []int, direction string, visited [][]int) {
	for row, _ := range area {
		for col, _ := range area[row] {
			if row == pos[0] && col == pos[1] {
				fmt.Print(direction)
			} else if area[row][col] == "^" {
				fmt.Print(".")
			} else {
				fmt.Print(area[row][col])
			}

		}
		fmt.Println()
	}
}
