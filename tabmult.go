package checkpoint

import "github.com/01-edu/z01"

func tabmult(nbr int) {

	for j, e := 1, '1'; e <= '9'; e, j = e+1, j+1 {
		z01.PrintRune(e)
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		printstr(itoa(nbr))
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		printstr(itoa(nbr * j))
		z01.PrintRune('\n')
	}
}

func itoa(i int) string {
	var neg bool
	var result string
	if i == 0 {
		return "0"
	}
	if i < 0 {
		neg = true
		i *= -1
	}
	for i > 0 {
		result = string(i%10+48) + result
		i /= 10
	}
	if neg {
		result = "-" + result
	}
	return result
}

func printstr(s string) {
	runes := []rune(s)
	for i := range s {
		z01.PrintRune(runes[i])
	}
}
