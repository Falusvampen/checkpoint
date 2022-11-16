package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		nbr, _ := strconv.Atoi(os.Args[1])
		Tabmult(nbr)
	}
}

func Tabmult(nbr int) {
	for i := 1; i <= 9; i++ {
		z01.PrintRune(rune(i + '0'))
		printstr(" x ")
		printstr(itoa(nbr))
		printstr(" = ")
		printstr(itoa(i * nbr))
		z01.PrintRune('\n')
	}
}

func printstr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func itoa(i int) string {
	result := ""

	if i == 0 {
		return "0"
	}

	for i > 0 {
		result = string(i%10+'0') + result
		i /= 10
	}
	if i < 0 {
		return "-" + itoa(-i)
	}
	return result
}
