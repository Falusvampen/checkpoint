package main

import (
	"checkpoint"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	checkpoint.PrintStr(lastWord(os.Args[1]))
	z01.PrintRune('\n')
}

func rmEndSpace(s string) string {
	for s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s
}

func lastWord(s string) string {
	s = rmEndSpace(s)
	result := ""
	for _, c := range s {
		if c != ' ' {
			result += string(c)
		} else {
			result = ""
		}
	}
	return result
}
