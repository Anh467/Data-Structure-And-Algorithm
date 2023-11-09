package main

import "fmt"

/*
	Given a string s, return the number of homogenous substrings of s. Since the answer may be too large, return it modulo 109 + 7.

	A string is homogenous if all the characters of the string are the same.

	A substring is a contiguous sequence of characters within a string.



	Example 1:

	Input: s = "abbcccaa"
	Output: 13
	Explanation: The homogenous substrings are listed as below:
	"a"   appears 3 times.
	"aa"  appears 1 time.
	"b"   appears 2 times.
	"bb"  appears 1 time.
	"c"   appears 3 times.
	"cc"  appears 2 times.
	"ccc" appears 1 time.
	3 + 1 + 2 + 1 + 3 + 2 + 1 = 13.
	Example 2:

	Input: s = "xy"
	Output: 2
	Explanation: The homogenous substrings are "x" and "y".
	Example 3:

	Input: s = "zzzzz"
	Output: 15
*/

/*
	STATUS: Accepted
	Runtime: 0ms Beats 100.00%of users with Go
	Memory: 6.41MB (Beats 51.85%of users with Go)
*/
func countHomogenous(s string) int {
	mod := 1000000007
	s = s + "$"
	sum := 0
	index := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			sum += (index * (index + 1)) / 2
			index = 1
			continue
		}
		index += 1
	}
	return int(sum % (mod))
}
func main() {
	//hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhdllljbbbbiiiiiiiiiipppaaaaaaaqmmsssssspkkkkkkkkkkkelggggggggggggggppoooozzzaaaaaaazzqqqqqqqqqqqqiiiioooooooopeeeeemccvvvrrrkkkkffffqabbbbbbbbqqqjjjjjjttiiiissssssffffffggggcccrrrrrjjjjjjjpawwwwrggggggggggllllllvvvvvviiyyyggxxxzggvvhvhhhhhlllllllllllllllllbrrsssssaddbbbbbbbbbbbbbbbbbbbbbbbboaaaaaaaqbbbbbbbnnnnnaaaaaaaaapqqqqqqqqqqqqqgglllwwwwwwwwwnnnnwwwwwiiiiilllwwhhhhhhhhhhhhhppppppppppkkkkqqqqqqqqqqqqqqqqqkxxxxxssiikkjjjjjkooooooooooottttttthhhhhhhhccccccccccceccmmrrrrrdddddddduuuuuuuuukkibbxxxxxxxxxxmmmqqqqqqqqqqqgggzzzzhhiiyyyyhhhhhhhxx
	s := []string{
		"abbcccxx",
		"zzzzzz",
		"hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhdllljbbbbiiiiiiiiiipppaaaaaaaqmmsssssspkkkkkkkkkkkelggggggggggggggppoooozzzaaaaaaazzqqqqqqqqqqqqiiiioooooooopeeeeemccvvvrrrkkkkffffqabbbbbbbbqqqjjjjjjttiiiissssssffffffggggcccrrrrrjjjjjjjpawwwwrggggggggggllllllvvvvvviiyyyggxxxzggvvhvhhhhhlllllllllllllllllbrrsssssaddbbbbbbbbbbbbbbbbbbbbbbbboaaaaaaaqbbbbbbbnnnnnaaaaaaaaapqqqqqqqqqqqqqgglllwwwwwwwwwnnnnwwwwwiiiiilllwwhhhhhhhhhhhhhppppppppppkkkkqqqqqqqqqqqqqqqqqkxxxxxssiikkjjjjjkooooooooooottttttthhhhhhhhccccccccccceccmmrrrrrdddddddduuuuuuuuukkibbxxxxxxxxxxmmmqqqqqqqqqqqgggzzzzhhiiyyyyhhhhhhhxx",
	}
	sum := countHomogenous(s[1])
	fmt.Println(sum)
}
