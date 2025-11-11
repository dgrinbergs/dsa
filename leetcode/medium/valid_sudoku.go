package medium

// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]int)
	cols := make(map[int]map[byte]int)
	squares := make(map[int]map[byte]int)

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if rows[row] == nil {
				rows[row] = make(map[byte]int)
			}

			if cols[col] == nil {
				cols[col] = make(map[byte]int)
			}

			square := ((row / 3) * 3) + (col / 3)
			if squares[square] == nil {
				squares[square] = make(map[byte]int)
			}

			num := board[row][col]

			if num == '.' {
				continue
			}

			if _, ok := squares[square][num]; ok {
				return false
			}

			if _, ok := rows[row][num]; ok {
				return false
			}

			if _, ok := cols[col][num]; ok {
				return false
			}

			rows[row][num]++
			cols[col][num]++
			squares[square][num]++
		}
	}

	return true
}
