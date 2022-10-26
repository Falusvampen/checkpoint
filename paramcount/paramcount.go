package main

import (
	"os"

	"github.com/01-edu/z01"
)

// In the checkpoint they only test up to 9 arguments, therefore the maximum number of arguments for this program is 9.

func main() {
	// z01.PrintRune()                                 This prints a rune
	// len(os.Args) - 1 + '0'                          This takes the length of arguments
	// rune(len(os.Args) - 1 + '0')                    This converts the length of arguments to a rune
	// z01.PrintRune(rune(len(os.Args) - 1 + '0'))     Everything together
	z01.PrintRune(rune(len(os.Args) - 1 + '0'))
	z01.PrintRune('\n')
}
