package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartracted") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}

	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}
