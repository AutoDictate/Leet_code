package main

func Zeros(n int) (count int) {

	for n >= 5 {
		n /= 5
		count += n
	}
	return

}