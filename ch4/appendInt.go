package main

import "fmt"

func main() {
	var numbers []int

	numbers = appendInt(numbers, 1)
	numbers = appendInt(numbers, 2)
	numbers = appendInt(numbers, 3)
	numbers = appendInt(numbers, 4)
	numbers = appendInt(numbers, 5)

	fmt.Println(numbers)
}

func appendInt(original []int, newItem int) []int {
	var result []int
	newLength := len(original) + 1

	if newLength <= cap(original) {
		result = original[:newLength]
	} else {
		newCap := newLength

		if newCap < 2*len(original) {
			newCap = 2 * len(original)
		}

		result = make([]int, newLength, newCap)
		copy(result, original)
	}

	result[len(original)] = newItem
	return result
}
