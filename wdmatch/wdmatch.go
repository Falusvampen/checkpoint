package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	i := 0
	j := 0
	for i < len(arg1) && j < len(arg2) {
		if arg1[i] == arg2[j] {
			i++
		}
		j++
	}
	if i == len(arg1) {
		// for _, c := range arg1 {
		// 	z01.PrintRune(c)
		// }
		for _, c := range arg1 {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
