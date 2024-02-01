package main

func intersect(nums1 []int, nums2 []int) []int {

	seen := make(map[int]int)

	for _, num := range nums1 {
		seen[num]++
	}

	result := []int{}

	for _, num := range nums2 {
		if seen[num] > 0 {
			result = append(result, num)
			seen[num]--
		}
	}

	return result

}
