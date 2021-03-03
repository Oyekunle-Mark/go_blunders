package main

func missingNumber(nums []int) int {
	numSet := make(map[int]bool)

	for _, num := range nums {
		numSet[num] = true
	}

	maxNum := len(nums)
	var missingNumber int

	for i := 0; i <= maxNum; i++ {
		if _, ok := numSet[i]; !ok {
			missingNumber = i
		}
	}

	return missingNumber
}
