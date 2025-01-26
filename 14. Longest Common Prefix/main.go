package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	first := strs[0]
	var answer string

	for i := 0; i < len(first); i++ {
		currentChar := first[i]

		for _, j := range strs {
			if i >= len(j) || j[i] != currentChar {
				return answer
			}
		}

		answer += string(currentChar)
	}

	return answer
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

	emptyStrs := []string{}
	fmt.Println(longestCommonPrefix(emptyStrs))
}
