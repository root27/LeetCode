package main

func firstUniqChar(s string) int {

	// Create a map to store the frequency of each character
	frequency := make(map[rune]int)

	// Iterate over the string and store the frequency of each character
	for _, char := range s {
		frequency[char]++
	}

	// Iterate over the string again and return the index of the first character with a frequency of 1
	for i, char := range s {
		if frequency[char] == 1 {
			return i
		}
	}

	// If no unique character is found, return -1
	return -1
}
