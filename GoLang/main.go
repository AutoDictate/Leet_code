package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// res := LongestSubsequence("abcabcbb")
	// fmt.Println(res)
	
	// res := findMedianSortedArrays([]int{1,3}, []int{2, 4})
	// fmt.Printf("%f", res)

	// res := LongestPalindrome("abcda")
	// fmt.Println(res)

	// res := RomanNumeralDecoder("MMMXLIX")
	// fmt.Println(res)

	// res:= ValidBraces("(){}[]")
	// fmt.Println(res)

	// res := Disemvowel("This website is for losers LOL!")
	// fmt.Println(res)

	// res := Accum("ZpglnRxqenU")
	// fmt.Println(res)

	// res := CamelCase("hello world")
	// fmt.Println(res)

	// res:= Order("4of Fo1r pe6ople g3ood th5e the2")
	// fmt.Println(res)


	res := LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},1 )
	fmt.Println(res)

}