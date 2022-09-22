package checkpoint

func intToRoman(num int) string {

	if num > 3999 {
		return "ohno"
	}

	symbols := []string{
		"M", "CM", "D", "CD", "C", "XC", "L",
		"XL", "X", "IX", "V", "IV",
		"I"}

	values := []int{
		1000, 900, 500, 400, 100,
		90, 50, 40, 10, 9, 5, 4,
		1}
	roman := ""
	i := 0
	for num > 0 {
		k := num / values[i]
		for j := 0; j < k; j++ {
			roman += symbols[i]
			num -= values[i]
		}
		i++
	}
	return roman
}
