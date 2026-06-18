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

func capitalize(word string) string {
    if len(word) == 0 {
        return word
    }
    return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func applyCaseModifiers(text string) string {
    words := strings.Split(text, " ")
    result := []string{}

    for i := 0; i < len(words); i++ {
        w := words[i]

        // Check if this word is a tag like [up], [low], [cap], [up,3], [low,2], [cap,6]
        if strings.HasPrefix(w, "[") && strings.HasSuffix(w, "]") {
            inner := w[1 : len(w)-1]
            parts := strings.Split(inner, ",")
            modifier := strings.TrimSpace(parts[0])
            n := 1
            if len(parts) == 2 {
                num, err := strconv.Atoi(strings.TrimSpace(parts[1]))
                if err == nil {
                    n = num
                }
            }

            // Apply to the last n words in result
            start := len(result) - n
            if start < 0 {
                start = 0
            }

            for j := start; j < len(result); j++ {
                switch modifier {
                case "up":
                    result[j] = strings.ToUpper(result[j])
                case "low":
                    result[j] = strings.ToLower(result[j])
                case "cap":
                    result[j] = capitalize(result[j])
                }
            }
            // Don't add the tag to result
        } else {
            result = append(result, w)
        }
    }

    return strings.Join(result, " ")
}

func main() {
    text := readInput()
    fmt.Println(applyCaseModifiers(text))
}