package main

func intToRoman(num int) string {
	// maps roman symbols to their integer values
	table := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	// orders roman symbols from largest to smallest
	symbolOrder := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""      // hold the symbols as they are deduced
	currentIndex := 0 // signifies the current symbol in the symbolOrder

	// iterate while num is greater than zero
	for num > 0 {
		// if current symbol's integer value can be extracted from number. That is, is num bigger than or equals current symbol's integer values
		remainder := num / table[symbolOrder[currentIndex]]

		// if so, add the symbol to result, and decrement the integer value from num
		if remainder > 0 {
			result += symbolOrder[currentIndex]
			num -= table[symbolOrder[currentIndex]]
		} else { // otherwise, move down the symbol to the next one
			currentIndex++
		}
	}

	return result
}
