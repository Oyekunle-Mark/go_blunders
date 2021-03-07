package main

import "strings"

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func minimumLengthEncoding(words []string) int {
	encoding := ""

	for _, word := range words {
		wordInEncoding := strings.Contains(encoding, word)

		if !wordInEncoding {
			encoding += word + "#"
		}
	}

	return len(encoding)
}
