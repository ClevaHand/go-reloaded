package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func convertHex(text string) string {
	words := strings.Split(text, " ")
	result := []string{}
	i := 0
	for i < len(words) {
		if words[i] == "[hex]" && len(result) > 0 {
			last := result[len(result)-1]
			val, err := strconv.ParseInt(last, 16, 64)
			if err == nil {
				result[len(result)-1] = fmt.Sprintf("%d", val)
			}
		} else {
			result = append(result, words[i])
		}
		i++
	}
	return strings.Join(result, " ")
}

func main() {
	text := readInput()
	fmt.Println(convertHex(text))
}
