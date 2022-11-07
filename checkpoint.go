package checkpoint


func Rot14(s string) string {
	res := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			res += string((r-'A'+14)%26 + 'A')
		} else if r >= 'a' && r <= 'z' {
			res += string((r-'a'+14)%26 + 'a')
		}
	}
	return res
}

func LastWord(s string) string {
	res := ""
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == ' ' && i != len(s)-1 {
			break
		}
		res = string(s[i]) + res
	}
	return res
}
