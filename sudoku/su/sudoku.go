package sudoku

func isSafe(board [][]int, row, col, num int) bool {
    for i := 0; i < 9; i++ {
        if board[row][i] == num || board[i][col] == num {
            return false
        }
    }

    startRow, startCol := row-row%3, col-col%3
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i+startRow][j+startCol] == num {
                return false
            }
        }
    }

    return true
}

func solveSudokuUtil(board [][]int) bool {
    empty := true
    var row, col int

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == 0 {
                row, col = i, j
                empty = false
                break
            }
        }
        if !empty {
            break
        }
    }

    if empty {
        return true
    }

    for num := 1; num <= 9; num++ {
        if isSafe(board, row, col, num) {
            board[row][col] = num
            if solveSudokuUtil(board) {
                return true
            }
            board[row][col] = 0
        }
    }

    return false
}

func SolveSudoku(board [][]int) [][]int {
    solveSudokuUtil(board)
    return board
}