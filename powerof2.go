package checkpoint

func isPowerOf2(n int) bool {
	return n != 0 && ((n & (n - 1)) == 0)
}
