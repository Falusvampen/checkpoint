package main

import (
	"checkpoint"
	"os"
)

func compare(s, s2 string) {
	// make a variable to store the results
	// loop through the first string
	// loop through the second string and remember to save the index
	result := ""
	for _, e := range s {
		
		for i, e2 := range s2 {
			if e == e2 {
				// if the characters match, add the character to the result, 
				result += string(e)
				s2 = s2[i:]
				break
			}
		}
	}
	if result == s {
		checkpoint.PrintStr(result)
	}
}

func main() {
	// make sure there are two arguments
	if len(os.Args) != 3 {
		return
	}
	compare(os.Args[1], os.Args[2])
}
