package main

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
