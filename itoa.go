package checkpoint

func Itoa(i int) string {

	if i == 0 {
		return "0"
	}
	res := ""
	for i > 0 {
		res = string((i%10 + 48)) + res
		i /= 10
	}
	if i < 0 {
		return "-" + Itoa(-i)
	} else {
		return res
	}
}

// func Itoa(n int) string {
// 	if n == 0 {
// 		return "0"
// 	}
// 	res := ""
// 	for n > 0 {
// 		res = string(n%10+48) + res
// 		n /= 10
// 	}
// 	if n < 0 {
// 		return "-" + Itoa(-n)
// 	} else {
// 		return res
// 	}
// }
