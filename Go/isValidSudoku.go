package main

func isValidSudoku(board [][]byte) bool {

	// Check rows
	for i := 0; i < 9; i++ {
		seen := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if seen[board[i][j]] {
					return false
				}
				seen[board[i][j]] = true
			}
		}
	}

	// Check columns
	for i := 0; i < 9; i++ {
		seen := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				if seen[board[j][i]] {
					return false
				}
				seen[board[j][i]] = true
			}
		}
	}

	// Check squares
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {

			seen := make(map[byte]bool)

			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					if board[k][l] != '.' {
						if seen[board[k][l]] {
							return false
						}
						seen[board[k][l]] = true
					}
				}
			}

		}
	}

	return true

}
