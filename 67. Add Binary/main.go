package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	b = fmt.Sprintf("%0*s", len(a), b)

	carry := 0
	result := ""

	for i := len(a) - 1; i >= 0; i-- {
		r := carry
		if a[i] == '1' {
			r++
		}
		if b[i] == '1' {
			r++
		}
		result = fmt.Sprintf("%d%s", r%2, result)
		carry = r / 2
	}

	if carry > 0 {
		result = fmt.Sprintf("%d%s", carry, result)
	}

	return result
}

func main() {
	a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	result := addBinary(a, b)
	fmt.Println(result)
}
