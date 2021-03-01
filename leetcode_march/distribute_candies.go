package main

func distributeCandies(candyType []int) int {
	candySet := make(map[int]bool) // act as a set of distinct candies
	var distinctCandies []int      // hold distinct candies
	maxCandies := len(candyType) / 2

	// for all candy type
	for _, candy := range candyType {
		// if current candy isn't in set
		if _, ok := candySet[candy]; !ok {
			distinctCandies = append(distinctCandies, candy) // add to distinct candies slice
			candySet[candy] = true                           // add to set
		}
	}

	// if distinct candies is less than the max, return len of distinc
	if len(distinctCandies) < maxCandies {
		return len(distinctCandies)
	}

	return maxCandies // otherwise, she can eat the max
}
