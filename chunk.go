package checkpoint

import "fmt"

func Chunk(a []int, ch int) {
	if ch <= 0 {
		return
	}
	var slice []int
	result := make([][]int, 0, len(a)/ch+1)
	for len(a) >= ch {
		slice, a = a[:ch], a[ch:]
		result = append(result, slice)
	}
	if len(a) > 0 {
		result = append(result, a)
	}
	fmt.Println(result)
}
