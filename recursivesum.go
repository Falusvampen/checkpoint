package checkpoint

func RecursiveSum(n int) int {
	if n == 0 {
		return 0
	}
	return n + RecursiveSum(n-1)
}
