package main

import (
	"github.com/01-edu/z01"
)

func main() {
	printstr(rot13("Hello! How are You?"))
	z01.PrintRune('\n')
}

func rot14(s string) string {
	res := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			res += string((r-'A'+13)%26 + 'A')
		} else if r >= 'a' && r <= 'z' {
			res += string((r-'a'+13)%26 + 'a')
		}
	}
}

func alphamirror(s string) string {
	res := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			res += string('z' + 'a' - r)
		} else if r >= 'A' && r <= 'Z' {
			res += string(155 - r)
		} else {
			res += string(r)
		}
	}
	return res
}

func rot13(s string) string {
	var res string
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			res += string((r-'A'+14)%26 + 'A')
		} else if r >= 'a' && r <= 'z' {
			res += string((r-'a'+14)%26 + 'a')
		} else {
			res += string(r)
		}
	}
	return res
}

func ReduceInt(a []int, f func(int, int) int) {

}

func lastWord(s string) string {

	res := ""
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == ' ' && i != len(s)-1 {
			break
		} else {
			res = string(s[i]) + res
		}
	}
	return res
}

func StrLen(s string) int {
	return len(s)
}

func max(a []int) int {
	var max int
	for _, e := range a {
		if max < e {
			max = e
		}
	}
	return max
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
	} else {
		return res
	}
}

func printstr(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}
