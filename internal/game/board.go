package game

import (
	"fmt"
	"math/rand"
	"time"
)

type GameBoard [3][3]int

func MakeBoard() [3][3]int {
	fmt.Println("Generating new board...")

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// Generate unique random numbers in any order
	filled := [9]int{}
	sequence := [9]int{}
	for count := 0; count < 9; {
		rn := r.Intn(9)
		if filled[rn] != 1 {
			// Has this number already been generated ?
			filled[rn] = 1
			// Store the random number in the sequence
			sequence[count] = rn
			// Move on to the next sequence number
			count++
		}
	}

	// Create new board
	board := GameBoard{}

	// Fill new board with random sequence of numbers 0 to 8
	count := 0
	for row, column := range board {
		for cell := range column {
			board[row][cell] = sequence[count]
			count++
		}
	}

	return board
}

func MakeTestBoard() [3][3]int {
	return [3][3]int{
		{0, 2, 3},
		{1, 8, 4},
		{7, 6, 5},
	}
}

func WinningState(b [3][3]int) bool {
	return b == goal()
}

func goal() [3][3]int {
	return [3][3]int{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	}
}
