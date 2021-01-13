package palindrome

import "unicode"

// IsPalindrome checks if a string is palindrome
// returns true if it is and false otherwise
// (unoptimised)
func IsPalindrome(s string) bool {
	var letters []rune

	for _, letter := range s {
		if unicode.IsLetter(letter) {
			letters = append(letters, unicode.ToLower(letter))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}

	return true
}
