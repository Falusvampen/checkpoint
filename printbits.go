package checkpoint

import "github.com/01-edu/z01"

// print bits of a byte
func PrintBits(octet byte) {
	for i := 7; i >= 0; i-- {
		if octet&(1<<uint(i)) != 0 {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
	}
}
