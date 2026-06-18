package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fixPunctuation(text string) string {
	words := strings.Fields(text)
	var result []string

	for i := 0; i < len(words); i++ {
		// If it's punctuation and we have a previous word, attach it
		if i > 0 && strings.ContainsAny(words[i], "!?,;:.") {
			result[len(result)-1] += words[i]
		} else {
			result = append(result, words[i])
		}
	}

	return strings.Join(result, " ")
}

func main() {
	// Reads input from terminal
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		fmt.Println(fixPunctuation(text))
	}
}
