package main

import "strings"

func CamelCase(s string) string {

	var res string
	s = strings.ToLower(s)
	for i:= 0; i < len(s); i++ {

		v := string(s[i])

		if i == 0 && v != " " {
			res += strings.ToUpper(v)
		} else if string(v)== " " {
			continue
		} else if string(s[i-1]) == " " {
			res += strings.ToUpper(v)
		} else {
			res += v
		}
	}

	return res

}