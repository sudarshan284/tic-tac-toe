package main

import "fmt"

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"", "", ""},
		[]string{"", "", ""},
		[]string{"", "", ""},
	}

	// The players take turns.
	playGame(board)
}

func displayBoard(board [][]string) {
	fmt.Println("Current board")
	for _, row := range board {
		fmt.Println(row)
	}
	fmt.Println()
}

func makeMove(board [][]string, row, col int, player string) bool {
	if row < 0 || row > 2 || col < 0 || col > 2 || board[row][col] != "" {
		return false
	}

	board[row][col] = player
	return true
}

func checkWinner(board [][]string) string {
	for i := 0; i < 3; i++ {
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}

		if board[0][i] != "" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}

	if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}

	if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}
	return ""
}

func playGame(board [][]string) {
	players := []string{"X", "O"}
	turn := 0
	moves := 0

	for {
		displayBoard(board)
		player := players[turn%2]
		fmt.Printf("Player %s,enter your move(row,column): ", player)
		var row, col int
		fmt.Scan(&row, &col)

		if makeMove(board, row, col, player) {
			moves++
			winner := checkWinner(board)
			if winner != "" {
				displayBoard(board)
				fmt.Printf("Player %s wins !\n", winner)
				break
			}

			if moves == 9 {
				displayBoard(board)
			}
			turn++
		} else {
			fmt.Println("Invalid move,TRy again")
		}
	}
}
