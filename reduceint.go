package checkpoint

import "strconv"

// func ReduceInt(a []int, f func(int, int) int) {
// 	firstNum := a[0]
// 	secondNum := a[1]
// 	result := f(firstNum, secondNum)
// 	PrintStr(strconv.Itoa(result))
// }

// shorter version
func ReduceInt(a []int, f func(int, int) int) {
	PrintStr(strconv.Itoa(f(a[0], a[1])))
}
