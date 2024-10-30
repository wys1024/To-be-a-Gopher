package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	find(nums, len(nums))

}

func find(nums []int, n int) {
	visit := make(map[int]int, 0)
	for _, k := range nums {
		visit[k]++
	}

	for k, v := range visit {
		if v == n/2 {
			fmt.Println(k)
		}
	}
}