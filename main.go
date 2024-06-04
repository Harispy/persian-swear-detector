package main

import (
	"bufio"
	"strings"
)

type SwearWordDetector struct {
	swearWords map[string]struct{}
}

func NewCustomSwearWordDetector(words []string) *SwearWordDetector {
	swearWords := make(map[string]struct{})
	for _, word := range words {
		swearWords[word] = struct{}{}
	}
	return &SwearWordDetector{swearWords: swearWords}
}

func (iwd *SwearWordDetector) ContainsIllegalWord(text string) bool {
	text = strings.Map(persianNormalizer, text)

	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		token := scanner.Text()
		if _, found := iwd.swearWords[token]; found {
			return true
		}
	}
	return false
}
