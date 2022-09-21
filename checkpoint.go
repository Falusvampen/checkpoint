package checkpoint

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func RecursiveSum(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + RecursiveSum(n-1)
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

func Rot14(s string) string {
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

func ReduceInt(a []int, f func(int, int) int) {

}

func LastWord(s string) string {

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

func Max(a []int) int {
	var max int
	for _, e := range a {
		if max < e {
			max = e
		}
	}
	return max
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for n > 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	if n < 0 {
		return "-" + Itoa(-n)
	} else {
		return res
	}
}

func PrintStr(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}
