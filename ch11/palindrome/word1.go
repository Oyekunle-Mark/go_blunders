package palindrome

// IsPalindrome checks if a string is palindrome
// returns true if it is and false otherwise
// (unoptimised)
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
