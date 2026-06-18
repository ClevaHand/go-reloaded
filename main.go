package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads all input from terminal
func readInput() string {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n")
}if words[i] == "(bin)"

// Converts words before (bin) from binary to decimal
func convertBin(text string) string {
	words := strings.Fields(text)

	var result []string

	i := 0

	for i < len(words) {

		if words[i] == "(bin)" && len(result) > 0 {

			// Get previous word
			last := result[len(result)-1]

			// Convert binary → decimal
			val, err := strconv.ParseInt(last, 2, 64)

			if err == nil {
				result[len(result)-1] = strconv.FormatInt(val, 10)
			}

		} else {
			result = append(result, words[i])
		}

		i++
	}

	return strings.Join(result, " ")
}

func main() {
	fmt.Println("Enter text and press Ctrl+D (Linux/macOS) or Ctrl+Z then Enter (Windows):")

	input := readInput()

	output := convertBin(input)

	fmt.Println("\nOutput:")
	fmt.Println(output)
}
