package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

func solveDay1(data []string) (int, int) {
	// Part One
	var answer1 float64 = 0

	lefts := make([]int, len(data))
	rights := make([]int, len(data))
	fmt.Println(len(data), len(lefts))
	for i, v := range data {
		pair := strings.Split(v, "   ")
		left := stringToInt(pair[0])
		right := stringToInt(pair[1])
		lefts[i] = left
		rights[i] = right
	}

	slices.Sort(lefts)
	slices.Reverse(lefts)
	slices.Sort(rights)
	slices.Reverse(rights)

	for i, _ := range lefts {
		answer1 += math.Abs(float64(lefts[i]) - float64(rights[i]))
	}

	// Part Two
	var answer2 int = 0
	for _, l := range lefts {
		count := 0
		for _, r := range rights {
			if r == l {
				count++
			}
		}
		answer2 += count * l
	}

	return int(answer1), answer2

}
