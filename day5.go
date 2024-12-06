package main

import (
	"fmt"
	"strings"
)

func solveDay5(data []string) (int, int) {
	answer1 := 0
	answer2 := 0

	PrintSlice(data)

	// orderMap[num1]: []nums2 -> num1 comes before all nums2
	orderMap := make(map[int][]int)
	gapIndex := 0

	for idx, line := range data {
		if line == "" {
			gapIndex = idx
			break
		}
		nums := strings.Split(line, "|")
		left := stringToInt(nums[0])
		right := stringToInt(nums[1])

		if orderMap[left] == nil {
			orderMap[left] = []int{right}
		} else {
			orderMap[left] = append(orderMap[left], right)
		}

	}
	fmt.Println(orderMap)

	pages := data[gapIndex+1:]

	fmt.Println(pages)

	for _, page := range pages {
		splitPage := Map(strings.Split(page, ","), func(item string) int { return stringToInt(item) })
		pageIndex := 0
		validPage := true
		for {
			if pageIndex > len(splitPage)-1 {
				break
			}
			num := splitPage[pageIndex]
			if orderMap[num] != nil {
				// has rules associated with it
				result := followsRules(splitPage, pageIndex, orderMap)
				if !result {
					validPage = false
				}
			}
			pageIndex++
		}

		if validPage {
			middle := splitPage[len(splitPage)/2]
			fmt.Println("Follows Rules:", splitPage, middle)
			answer1 += middle
		} else {
			fixed := fixOrdering(splitPage, orderMap)
			middle := fixed[len(fixed)/2]
			answer2 += middle
		}
	}

	return answer1, answer2

}

func followsRules(page []int, index int, orderMap map[int][]int) bool {
	for _, p := range page[:index] {
		for _, invalidNum := range orderMap[page[index]] {
			if p == invalidNum {
				return false
			}
		}
	}
	return true
}

// Given a page and the rules, returns a correct ordering
func fixOrdering(page []int, rules map[int][]int) []int {

	fmt.Println("Before Fix:", page)
	for pi, _ := range page {
		numsBefore := rules[page[pi]]
		for i, v := range page[:pi] {
			for _, n := range numsBefore {
				if v == n {
					// problem, swap
					temp := page[i]
					page[i] = page[pi]
					page[pi] = temp
				}
			}
		}
	}
	fmt.Println("Fixed Ordering?", page)

	return page
}
