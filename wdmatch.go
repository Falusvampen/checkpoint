package main

import (
	"os"
)
2147483647
func main() {
	if len(os.Args) == 3 {
		arg1 := os.Args[1]
		arg2 := os.Args[2]
		if len(arg1) > len(arg2) {
			return
		}
		res := ""
		for _, e := range arg1 {
			for i, e2 := range arg2 {
				if e == e2 {
					res += string(e)
				}
			}
		}
	}
}
