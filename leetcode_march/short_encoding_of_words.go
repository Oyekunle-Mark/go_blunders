package main

import "strings"

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
