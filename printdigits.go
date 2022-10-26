package checkpoint

import "github.com/01-edu/z01"

func PrintDigits() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
