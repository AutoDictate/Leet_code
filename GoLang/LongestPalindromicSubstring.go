package main


func LongestPalindrome(s string) string {

	max := 0

	result := ""

	for i := 0; i < len(s); i++ {
		res, err := SubString(i, s, "")
		if len(res) > max && len(err) == 0 {
			max = len(res)
			result = res
		}
	}

	return result

}

func SubString(n int, s string, val string) (string, string) {

	if n == len(s) {
		return val, "nope"
	}
	val += string(s[n])
	
	n++
	
	if CheckPalindrome(val) && len(val)>1 {
		return val, ""
	}
	
	return SubString(n, s, val)

}

func CheckPalindrome(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
