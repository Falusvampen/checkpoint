package main

import (
	"checkpoint"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	checkpoint.PrintStr(lastWord(os.Args[1]))
}

func lastWord(s string) string {
	result := ""
	wordCount := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			wordCount++
		} else if wordCount > 0 {
			if s[i] == ' ' {
				break
			}
			result += string(s[i])
			// function to reverse a string
		}
	}
	return reverse(result)
}
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
