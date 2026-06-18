package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func readInput() string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return strings.Join(lines, "\n")
}

func fixSingleQuotes(text string) string {
	// Remove space after opening single quote
	re1 := regexp.MustCompile(`'\s+`)
	text = re1.ReplaceAllStringFunc(text, func(m string) string {
		return "'"
	})
	// Remove space before closing single quote
	re2 := regexp.MustCompile(`\s+'`)
	text = re2.ReplaceAllStringFunc(text, func(m string) string {
		return "'"
	})
	return text
}

func main() {
	text := readInput()
	fmt.Println(fixSingleQuotes(text))
}
