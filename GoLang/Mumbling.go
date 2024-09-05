package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {

	var res string

	s = strings.ToLower(s)

	for i, v := range s {
		
		d := i

		fmt.Println("Value of d :", d)
		t := strings.ToUpper(string(v))

		res += t + string(d)
		
		for d > 0 {
			d--
			res += string(v)
		}
		if i < len(s)-1{
			res += "-"
		}
	}

	return res
}
