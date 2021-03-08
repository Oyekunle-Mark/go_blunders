package main

func removePalindromeSub(s string) int {
	// return 0 for empty string
	if len(s) == 0 {
		return 0
	}

	// if string is palindrome, return 1
	if isPalindrome(s) {
		return 1
	} else { // otherwise, two
		return 2
	}
}

// isPalindrome returns true if s is a palindrome and false otherwise.
func isPalindrome(s string) bool {
	for i, c := range s {
		if uint8(c) != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
