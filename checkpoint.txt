package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	fmt.Println(compare("Hello!", "Hello!"))
	fmt.Println(compare("Salut!", "lut!"))
	fmt.Println(compare("Ola!", "Ol"))
}

func compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func recursivesum(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + recursivesum(n-1)
	}
}

func Chunk(slice []int, size int) {

	if size <= 0 {
		fmt.Println("")
	} else if len(slice) == 0 {
		fmt.Println(slice)
	} else {
		a := make([][]int, 0, size)
		for size < len(slice) {
			a = append(a, slice[0:size])
			slice = slice[size:]
		}
		fmt.Println(append(a, slice))
	}
}

func rot14(s string) string {
	res := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			res += string((r-'A'+14)%26 + 'A')
		} else if r >= 'a' && r <= 'z' {
			res += string((r-'a'+14)%26 + 'a')
		}
	}
	return res
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
			res += string((r-'A'+13)%26 + 'A')
		} else if r >= 'a' && r <= 'z' {
			res += string((r-'a'+13)%26 + 'a')
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
