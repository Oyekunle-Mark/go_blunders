package main

func main() {

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

	result[newLength] = newItem
	return result
}
