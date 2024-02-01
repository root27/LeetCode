package main

func moveZeroes(nums []int) {

	if len(nums) == 0 {
		return
	}

	i := 0

	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}

	for i < len(nums) {
		nums[i] = 0
		i++
	}

}
