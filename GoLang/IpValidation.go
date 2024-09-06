package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {

	ipv4 := strings.Split(ip, ".")

	if len(ipv4) < 4 || len(ipv4) > 4 {
		return false
	}

	for i:=0; i<len(ipv4); i++ {

		t:= ipv4[i]

		if string(t[0]) == "0" && len(t) > 1 {
			return false
		}

		ipVal, err := strconv.Atoi(ipv4[i])

		fmt.Println(ipVal)

		if err != nil {
			return false
		}

		if ipVal > 255 || ipVal < 0 {
			return false
		}
	}

	return true

}