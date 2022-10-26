package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Display last parameter
func main() {
	printstr(os.Args[len(os.Args)-1])
	z01.PrintRune('\n')
}

func printstr(s string) {
	for _, e := range s {
		z01.PrintRune(e)
	}
}
