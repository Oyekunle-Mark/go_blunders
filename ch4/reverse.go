package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	reverseInt(numbers)

	fmt.Println(numbers)
}

func reverseInt(original []int) {
	for i, j := 0, len(original)-1; i < j; i, j = i+1, j-1 {
		original[i], original[j] = original[j], original[i]
	}
}
