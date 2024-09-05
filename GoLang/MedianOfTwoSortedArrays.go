package main

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)

	l := len(nums1) / 2

	if len(nums1) % 2 != 0 {
		return float64(nums1[l])
	}

	return float64(float64(nums1[l]+nums1[l-1]) / float64(2))

}
