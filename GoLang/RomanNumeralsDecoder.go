package main


/*
Symbol    Value
I          1
V          5
X          10
L          50
C          100
D          500
M          1,000
*/

func RomanNumeralDecoder(s string) int {

	sv := make(map[string]int)

	sv["I"] = 1
	sv["V"] = 5
	sv["X"] = 10
	sv["L"] = 50
	sv["C"] = 100
	sv["D"] = 500
	sv["M"] = 1000

	sum := 0
	curValue := 0
	preValue := 0

	for i := len(s)-1; i >= 0; i-- {
		
		curValue = sv[string(s[i])]

		if curValue < preValue {
			sum -= curValue
		} else {
			sum += curValue
		}
		preValue = curValue
		
	}

	return sum

}