package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	strconv.Atoi(os.Args[1])
	if len(os.Args) == 4 {
		num1, err2 := atoi(os.Args[1])
		mod := os.Args[2]
		num2, err1 := atoi(os.Args[3])

		if err1 == "Error" || err2 == "Error" {
			return
		}
		number, err := numresult(num1, num2, mod)
		if err != "Error" {
			printstr(itoa(number))
		}
	}
}

func printstr(s string) {
	for _, e := range s {
		z01.PrintRune(e)
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for n > 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	if n < 0 {
		return "-" + itoa(-n)
	}
	return res
}

func numresult(i, j int, s string) (int, string) {
	res := 0
	err := ""
	if len(s) == 1 {
		switch s {
		case "+":
			res, err = (i + j), ""
		case "-":
			res, err = i-j, ""
		case "*":
			res, err = i*j, ""
		case "/":
			if i == 0 || j == 0 {
				return 0, "No division by 0"
			}
			res, err = i/j, ""
		case "%":
			if i == 0 || j == 0 {
				return 0, "No modulo by 0"
			}
			res, err = i%j, ""
		default:
			return 0, ""
		}
	}
	if res <= 9223372036854775807 && res >= -9223372036854775807 {
		return res, err
	} else {
		return 0, "Error"
	}
}

func atoi(s string) (int, string) {
	var res int
	var err string
	var sign int

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else {
		sign = 1
	}
	for _, v := range s {
		if v >= '0' && v <= '9' {
			res = res*10 + int(v-'0')
		} else {
			err = "Error"
		}
	}
	return res * sign, err
}
