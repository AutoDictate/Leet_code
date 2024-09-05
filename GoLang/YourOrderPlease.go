package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Order(s string) string {

	if len(s) == 0 {
		return s
	}

	res := strings.Split(s, " ")
	re := regexp.MustCompile("[0-9]+")

	 r := make([]string, len(res) + 1)


	var t int

	for _, v := range res {

		i := re.FindAllString(v, -1)
		t,_ = strconv.Atoi(i[0])
		r[t] = v
		
	}

	return strings.TrimSpace(strings.Join(r, " "))
}