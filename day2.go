package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

func checkLevels(levels []int) bool {
	var prev int
	var direction string
	for _, v := range levels {
		if prev == 0 {
			fmt.Println("No prev")
			prev = v
			continue
		}
		isSafe, dir := checkSafeSequence(prev, v, direction)
		direction = dir
		if !isSafe {
			return false
		}
		prev = v
	}
	return true
}

func checkLevels2(levels []int) bool {
	var prev int
	var direction string
	for _, v := range levels {
		if prev == 0 {
			fmt.Println("No prev")
			prev = v
			continue
		}
		isSafe, dir := checkSafeSequence(prev, v, direction)
		direction = dir
		if !isSafe {
			for idx := range len(levels) {
				newLevels := slices.Concat(levels[:idx], levels[idx+1:])
				success := checkLevels(newLevels)
				if success {
					return true
				}
			}
			return false
		}
		prev = v
	}
	return true
}

func checkSafeSequence(prev int, curr int, direction string) (bool, string) {
	diff := math.Abs(float64(prev) - float64(curr))
	if diff < 1 || diff > 3 {
		fmt.Println("difference  issue")
		return false, direction
	}
	if direction == "" {
		fmt.Println("no direction")
		if curr > prev {
			fmt.Println("setting dir up")
			direction = "up"
		} else {
			fmt.Println("setting dir down")
			direction = "down"
		}
	} else if direction == "up" && curr < prev {
		fmt.Println("wrong direction up")
		return false, direction
	} else if direction == "down" && curr > prev {
		fmt.Println("wrong direction down")
		return false, direction
	}

	return true, direction
}

func solveDay2(data []string) (int, int) {
	// Part One
	var answer1 int = 0
	var answer2 int = 0
	PrintSlice(data)
	for _, report := range data {
		levels := strings.Split(report, " ")
		isSafe := checkLevels(Map(levels, func(item string) int { return stringToInt(item) }))
		if isSafe {
			answer1++
		}

		isSafe2 := checkLevels2(Map(levels, func(item string) int { return stringToInt(item) }))
		if isSafe2 {
			answer2++
		}
	}

	return answer1, answer2
}
