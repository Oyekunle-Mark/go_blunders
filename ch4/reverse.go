package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	reverseIntInPlace(numbers)

	fmt.Println(numbers)
	fmt.Println(reverseInt(numbers))
	fmt.Println(numbers)
}

// reverseIntInPlace reverses a slice of int in place
func reverseIntInPlace(original []int) {
	for i, j := 0, len(original)-1; i < j; i, j = i+1, j-1 {
		original[i], original[j] = original[j], original[i]
	}
}

// reverseInt returns a reverse of the input slice
func reverseInt(original []int) []int {
	reverse := make([]int, len(original), cap(original))
	reverseIndex := len(original) - 1

	for _, v := range original {
		reverse[reverseIndex] = v
		reverseIndex--
	}

	return reverse
}
