package checkpoint

func Alphamirror(s string) string {
	res := ""
	for _, e := range s {
		if e >= 'A' && e <= 'Z' {
			res += string('A' - e + 'Z')
		} else if e >= 'a' && e <= 'z' {
			res += string('a' - e + 'z')
		} else {
			res += string(e)
		}
	}
	return res
}
