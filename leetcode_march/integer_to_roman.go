package main

func intToRoman(num int) string {
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

	numeralsOrder := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	currentIndex := 0

	for num > 0 {
		remainder := num / table[numeralsOrder[currentIndex]]

		if remainder > 0 {
			result += numeralsOrder[currentIndex]
			num -= table[numeralsOrder[currentIndex]]
		} else {
			currentIndex++
		}
	}

	return result
}
