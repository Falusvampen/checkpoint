package checkpoint

func Alphamirror(s string) string {
	res := ""
	for _, e := range s {
		if e >= 'A' && e <= 'Z' {
			res += string('A' + 'Z' - e)
		} else if e >= 'a' && e <= 'z' {
			res += string('a' + 'z' - e)
		} else {
			res += string(e)
		}
	}
	return res
}
