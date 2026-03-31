package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

// ["flower","flow","flight"]
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return strs[0]
	}
	base := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(base) > 0 {
			if strings.HasPrefix(strs[i], base) {
				break
			}
			base = base[:len(base)-1]
		}
	}
	return base
}
