package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Display last parameter
func main() {

	args := os.Args[len(os.Args)-1]
	for _, e := range args {
		z01.PrintRune(e)
	}
	z01.PrintRune('\n')
}
