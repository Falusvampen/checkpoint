package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1))
}

func intToRoman(num int) string {
	// Define a map of Roman numerals and their corresponding decimal values
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	// Define a slice of decimal values, in descending order
	decimals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result string // result is the Roman numeral representation of the input integer

	// Loop through the decimal values
	for _, decimal := range decimals {
		for num >= decimal { // If the input integer is greater than or equal to the decimal value
			result += romanMap[decimal] // Add the corresponding Roman numeral to the result
			num -= decimal              // Subtract the decimal value from the input integer
		}
	}

	return result
}
