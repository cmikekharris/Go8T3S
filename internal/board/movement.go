package board

import (
	"fmt"
)

func findEmptySpace(board [3][3]int) (int, int) {
	fmt.Println("Finding empty space...")

	for row, column := range board {
		for cell := range column {
			if board[row][cell] == 0 {
				return row, cell
			}
		}
	}

	return -1, -1
}

func move(board [3][3]int, direction string) [3][3]int {

	fmt.Println("Move", direction, "...")

	newBoard := board
	oldRow, oldCell := findEmptySpace(board)
	newRow, newCell := oldRow, oldCell

	switch direction {
	case "up":
		newRow = oldRow - 1
	case "down":
		newRow = oldRow + 1
	case "left":
		newCell = oldCell - 1
	case "right":
		newCell = oldCell + 1
	}

	// Return existing board if out of bounds
	if newRow < 0 || newRow > 2 || newCell < 0 || newCell > 2 {
		return board
	}

	// Swap empty space with number
	oldValue := board[newRow][newCell]
	newBoard[newRow][newCell] = 0
	newBoard[oldRow][oldCell] = oldValue

	return newBoard
}
