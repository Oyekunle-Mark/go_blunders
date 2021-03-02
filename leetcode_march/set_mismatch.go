package main

import "sort"

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	var result []int
	currentNum := 1

	for i, num := range nums {
		if i > 0 {
			if num == nums[i-1] {
				result = append(result, num)

				if num == currentNum {
					result = append(result, currentNum-1)
				} else {
					result = append(result, currentNum)
				}

				return result
			}
		}

		currentNum++
	}

	return result
}
