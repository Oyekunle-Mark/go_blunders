package main

func missingNumber(nums []int) int {
	var numsSum int     // hold the sum of integers in num
	maxNum := len(nums) // maximum number

	// find numsSum
	for i := 0; i <= maxNum; i++ {
		numsSum += i
	}

	// subtract every integer in nums from numsSum
	for _, num := range nums {
		numsSum -= num
	}

	return numsSum // remainder is the missing number
}
