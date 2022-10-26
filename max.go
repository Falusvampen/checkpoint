package checkpoint

func Max(a []int) int {
	res := 0
	for _, e := range a {
		if res < e {
			res = e
		}
	}
	return res
}
