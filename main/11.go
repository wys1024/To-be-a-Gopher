package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
	nums1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	merge(nums1, len(nums1), nums2, len(nums2))

	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:n], nums2...)
	sort.Ints(nums1)
}
