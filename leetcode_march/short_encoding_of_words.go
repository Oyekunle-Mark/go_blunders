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
	encoding := ""

	for _, word := range words {
		subStringStartIndex := strings.Index(encoding, word)

		if subStringStartIndex != -1 && encoding[subStringStartIndex + len(word)] == '#' {
			encoding += word + "#"
		}
	}

	return len(encoding)
}
