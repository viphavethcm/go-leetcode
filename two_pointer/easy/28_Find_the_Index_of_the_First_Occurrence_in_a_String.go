package two_pointer_easy

/*
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description
*/

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
