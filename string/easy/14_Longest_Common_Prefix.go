package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	base := strs[0]
	for _, str := range strs[1:] {
		for len(base) > 0 {
			if strings.HasPrefix(str, base) {
				break
			}
			base = base[0 : len(base)-1]
		}
	}
	return base
}
