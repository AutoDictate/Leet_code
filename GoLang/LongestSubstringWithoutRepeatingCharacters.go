package main

func LongestSubsequence(s string) int {
	
	max := 1

	for i := range s {
		t := CheckLong(i, s)
		if t > max {
			max = t
		}
	}

	return max

}

func CheckLong(n int, s string) int {

	t := 0
	newMap := make(map[string]int)

	for i:= n; i< len(s); i++ {

		str := string(s[i])

		if newMap[str] == 0 {
			newMap[str]++
			t +=1
		} else {
			break
		}
	}

	return t
}