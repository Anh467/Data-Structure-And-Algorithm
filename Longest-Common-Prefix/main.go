package main

import "fmt"

/*
	STATUS: Accepted
	Runtime: 0ms (Beats 100.00%of users with Go)
	Memory: 2.36MB (Beats 42.67%of users with Go)
*/
/*
	Example 1:

	Input: strs = ["flower","flow","flight"]
	Output: "fl"
	Example 2:

	Input: strs = ["dog","racecar","car"]
	Output: ""
	Explanation: There is no common prefix among the input strings.
*/
func longestCommonPrefix(strs []string) string {
	strFirst := strs[0]
	index := len(strFirst)

	for i := 1; i < len(strs); i++ {
		if index > len(strs[i]) {
			index = len(strs[i])
		}
		for j := 0; j < index; j++ {

			if strFirst[j] != strs[i][j] {
				index = j
			}
		}
	}
	if index < 0 {
		return ""
	}
	substring := strFirst[0:index]
	return substring
}

func main() {
	//"flower","flow","flight"
	// "a"
	// ["a","b"]
	strs := []string{"a", "b"}
	s := longestCommonPrefix(strs)
	fmt.Print(s)
}
