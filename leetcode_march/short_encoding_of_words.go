package main

import (
	"sort"
	"strings"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func minimumLengthEncoding(words []string) int {
	sort.Sort(ByLength(words)) // sort words by length
	encoding := ""             // hold encoding

	// loop for all word in words
	for _, word := range words {
		subStringStartIndex := strings.LastIndex(encoding, word) // find the start index of the last substring of word in encoding, if word is substring

		// word is not a substring or substring in word is not immediately followed by an '#'
		if subStringStartIndex == -1 || encoding[subStringStartIndex+len(word)] != '#' {
			encoding += word + "#"
		}
	}

	return len(encoding) // return encoding length
}
