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
