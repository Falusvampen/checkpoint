package checkpoint

func Rot13(s string) string {
	res := ""
	for _, e := range s {
		if e >= 'A' && e <= 'Z' {
			res += string((e-'A'+13)%26 + 'A')
		} else if e >= 'a' && e <= 'z' {
			res += string((e-'a'+13)%26 + 'a')
		} else {
			res += string(e)
		}
	}
	return res
}
