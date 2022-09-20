package main

import "os"

func main() {

	if len(os.Args) == 4 {
		num1 := os.Args[1]
		mod := os.Args[2]
		num2 := os.Args[3]

		switch mod {
		case "+":
			println(num1 + num2)
		case "-":
			println(num1 - num2)
		case "*":
			println(num1 * num2)
		case "/":
			println(num1 / num2)
		default:
			println("Invalid operator")

		}
	}
}

func atoi(s string) int {
	var n int
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}
