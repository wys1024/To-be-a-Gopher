package main

func searchInsert(nums []int, target int) int {
	return lowerBound(nums, target)
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := (left + right) / 2
	for left < right {
		if nums[mid] < target {
			left = mid + 1
			lowerBound(left, right)
		} else {
			right = mid - 1
			lowerBound(left, right)
		}
	}
	return left

}
