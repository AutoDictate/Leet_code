package main

func PartsSums(ls []uint64) []uint64 {
	n := len(ls)
	if n == 0 {
		return []uint64{0}
	}

	res := make([]uint64, n+1)

	totalSum := uint64(0)
	for _, v := range ls {
		totalSum += v
	}

	for i := 0; i < n; i++ {
		res[i] = totalSum
		totalSum -= ls[i]
	}

	res[n] = 0

	return res
}
