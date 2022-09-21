package checkpoint

// converts string to int and also returns true if there is an error
func Atoi(s string) (int, bool) {
	sign := 1
	res := 0
	err := false

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for _, e := range s {
		if e >= '0' && e <= '9' {
			res = res*10 + int(e-48)
		} else {
			err = true
		}
	}
	return res * sign, err
}
