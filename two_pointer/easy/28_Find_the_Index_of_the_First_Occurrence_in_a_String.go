package main

import "fmt"

/*
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description
*/
func main() {
	haystack := "mississippi"
	needle := "issi"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	i := 0
	for i+len(needle) <= len(haystack) {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
		i++
	}
	return -1
}
