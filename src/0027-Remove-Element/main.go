package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}
