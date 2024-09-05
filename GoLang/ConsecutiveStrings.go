package main

import (
	"strings"
)

func LongestConsec(strarr []string, k int) string {

	fin := ""
	max := 0

	
	for i:=0; i <= len(strarr)-k; i++ {

		if k<=0 || k > len(strarr) {
			break
		}

		str := ""

		 // "tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"

		str = strings.Join(strarr[i:i+k],"")
		if len(str) > max {
			max = len(str)
			fin = str
		}
	}

	return fin

}