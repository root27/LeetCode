package main

func canJump(nums []int) bool {

	maxJump := 0
	for i, num := range nums {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+num)
	}
	return true

}
