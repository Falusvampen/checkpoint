package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		printstr(os.Args[1])
	}
}

func printstr(s string) {
	for _, e := range s {
		z01.PrintRune(e)
	}
}
