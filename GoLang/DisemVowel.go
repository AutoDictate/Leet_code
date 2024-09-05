package main

import "strings"

func Disemvowel(s string) string {

	var res strings.Builder

	for _, val := range s {
		if strings.ContainsRune("aeiouAEIOU", val) {
			continue
		}
		res.WriteRune(val)
	}
	return res.String()
}