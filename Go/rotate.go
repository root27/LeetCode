package main

import "log"

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 || k <= 0 {
		return
	}

	k %= n // In case k is greater than the length of the slice

	reverse := func(arr []int, start, end int) {
		for start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}

	// Reverse the entire array
	reverse(nums, 0, n-1)

	// Reverse the first k elements
	reverse(nums, 0, k-1)

	// Reverse the remaining elements
	reverse(nums, k, n-1)

}

// Test

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)

	log.Println(nums)
}
