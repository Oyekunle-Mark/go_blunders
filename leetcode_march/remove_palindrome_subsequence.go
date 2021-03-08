package main

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}

	if isPalindrome(s) {
		return 1
	} else {
		return 2
	}
}

func isPalindrome(s string) bool {
	for i, c := range s {
		if uint8(c) != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
