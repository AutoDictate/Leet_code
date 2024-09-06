package main

import "strings"

func Duplicate_count(s1 string) int {

	s1 = strings.ToLower(s1)

	newMap := make(map[string]int)
	count := 0

	for _, val := range s1 {
		newMap[string(val)]++
	}

	for _, v := range newMap {
		if v > 1 {
			count++
		}
	}

	return count
}
