package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var table [26]int
	for i := 0; i < len(s); i++ {
		table[int(s[i]-'a')]++
	}
	for i := 0; i < len(t); i++ {
		table[int(t[i]-'a')]--
		if table[int(t[i]-'a')] < 0 {
			return false
		}
	}
	return true
}
