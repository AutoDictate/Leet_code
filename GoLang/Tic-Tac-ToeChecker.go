package main

func IsSolved(board [3][3]int) int {

	player1 := CheckRow(board, 1) || CheckCol(board, 1) || CheckDiagonal(board, 1)
	player2 := CheckRow(board, 2) || CheckCol(board, 2) || CheckDiagonal(board, 2)

	if player1 {
		return 1
	} else if player2 {
		return 2
	} else if CheckEmpty(board) {
		return -1
	}

	return 0

}

func CheckEmpty(board [3][3]int) bool {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false

}

func CheckRow(board [3][3]int, p int) bool {

	for i := 0; i < 3; i++ {
		c := 0
		for j := 0; j < 3; j++ {
			if board[i][j] == p {
				c++
			}
		}
		if c == 3 {
			return true
		}
	}
	return false

}

func CheckCol(board [3][3]int, p int) bool {

	for i := 0; i < 3; i++ {
		c := 0
		for j := 0; j < 3; j++ {
			if board[j][i] == p {
				c++
			}
		}
		if c == 3 {
			return true
		}
	}
	return false

}

func CheckDiagonal(board [3][3]int, p int) bool {

	mainDiagonal := true
    for i := 0; i < 3; i++ {
        if board[i][i] != p {
            mainDiagonal = false
            break
        }
    }

    // Check the anti-diagonal
    antiDiagonal := true
    for i := 0; i < 3; i++ {
        if board[i][2-i] != p {
            antiDiagonal = false
            break
        }
    }

    return mainDiagonal || antiDiagonal
}