package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := ""
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for _, s := range strs {
			if i >= len(s) || s[i] != ch {
				return prefix
			}
		}
		prefix += string(ch)
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Longest common prefix:", longestCommonPrefix(strs))
}
