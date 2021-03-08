package main

func removePalindromeSub(s string) int {

}

func isPalindrome(s string) bool {
	for i, c := range s {
		if uint8(c) != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
