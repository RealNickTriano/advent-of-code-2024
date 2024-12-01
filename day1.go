package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func mapFunc(val string) []string {
	return strings.Split(val, " ")
}

func filterFunc(val []string) []string {
	s := make([]string, 2)
	s[0] = val[0]
	for _, v := range val[1:] {
		if v != " " {
			s[1] = v
		}

	}
	return s
}

func solveDay1(data []string) int {
	PrintSlice(data)
	answer := 0

	s := make([][]string, len(data))
	for i, v := range data {
		s[i] = mapFunc(v)
	}

	for i, v := range s {
		s[i] = filterFunc(v)
	}
	lefts := make([]int, 0)
	rights := make([]int, 0)
	for _, v := range s {
		l, _ := strconv.Atoi(v[0])
		r, _ := strconv.Atoi(v[1])
		lefts = append(lefts, l)
		rights = append(rights, r)
	}

	slices.Sort(lefts)
	slices.Sort(rights)
	slices.Reverse(lefts)
	slices.Reverse(rights)

	final := make([][]int, 0)
	for i, _ := range lefts {
		final = append(final, []int{lefts[i], rights[i]})
	}
	PrintSlice(final)
	for _, v := range final {
		left := v[0]
		right := v[1]
		diff := math.Abs(float64(left) - float64(right))
		answer += int(diff)
	}

	answer = 0

	for _, v := range lefts {
		count := 0
		for _, val := range rights {
			if val == v {
				count++
			}
		}
		answer += count * v
	}

	// PrintSlice(s)
	return answer
}
