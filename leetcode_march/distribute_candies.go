package main

func distributeCandies(candyType []int) int {
	candySet := make(map[int]bool)
	var distinctCandies []int
	maxCandies := len(candyType) / 2

	for _, candy := range candyType {
		if _, ok := candySet[candy]; !ok {
			distinctCandies = append(distinctCandies, candy)
			candySet[candy] = true
		}
	}

	if len(distinctCandies) < maxCandies {
		return len(distinctCandies)
	}

	return maxCandies
}
