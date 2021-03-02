package main

func findErrorNums(nums []int) []int {
	maxNum := len(nums)          // the max number in nums
	var result []int             // hold the result
	numMap := make(map[int]bool) // reflect existing numbers

	for _, num := range nums {
		if _, ok := numMap[num]; !ok { // if num isn't in numMap
			numMap[num] = true // set to true
		} else {
			result = append(result, num) // otherwise, it's the repeated number
		}
	}

	// loop for all distinct integers from 1 to maxNum
	for i := 1; i <= maxNum; i++ {
		if _, ok := numMap[i]; !ok { // of current number isn't in numMap, it's the missing number
			result = append(result, i) // insert into result
			break
		}
	}

	return result
}
