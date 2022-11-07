package main

import (
	"checkpoint"
	"strconv"
)

// import (
// 	"checkpoint"
// 	"strconv"
// )

// func main() {
// 	mul := func(acc int, cur int) int {
// 		return acc * cur
// 	}
// 	sum := func(acc int, cur int) int {
// 		return acc + cur
// 	}
// 	div := func(acc int, cur int) int {
// 		return acc / cur
// 	}

// 	as := []int{500, 2}
// 	ReduceInt(as, mul)
// 	ReduceInt(as, sum)
// 	ReduceInt(as, div)
// }

// func multiply(a int, b int) int {
// 	return a * b
// }

func ReduceInt(a []int, f func(int, int) int) {

	// 	// this is 500
	// 	firstNum := a[0]
	// 	// this is 2
	// 	secondNum := a[1]

	// 	result := f(firstNum, secondNum)
	// 	// convert int to string
	// 	strconv.Itoa(result)
	// 	// print the result
	// 	checkpoint.PrintStr(strconv.Itoa(result))
	// 	// everything together

	checkpoint.PrintStr(strconv.Itoa(f(a[0], a[1])))
}
