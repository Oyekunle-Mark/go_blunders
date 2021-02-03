package main

func hammingWeight(num uint32) int {
	var oneCount = 0

	for current := num; current > 0; current = current / 2 {
		rem := current % 2

		if rem == 1 {
			oneCount++
		}
	}

	return oneCount
}
