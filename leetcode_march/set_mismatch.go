package main

func findErrorNums(nums []int) []int {
	maxNum := len(nums)
	var result []int
	numMap := make(map[int]bool)

	for _, num := range nums {
		if _, ok := numMap[num]; !ok {
			numMap[num] = true
		} else {
			result = append(result, num)
		}
	}

	for i := 1; i <= maxNum; i++ {
		if _, ok := numMap[i]; !ok {
			result = append(result, i)
			break
		}
	}

	return result
}
