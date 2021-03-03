package main

func missingNumber(nums []int) int {
	var numsSum int
	maxNum := len(nums)

	for i := 0; i <= maxNum; i++ {
		numsSum += i
	}

	for _, num := range nums {
		numsSum -= num
	}

	return numsSum
}
