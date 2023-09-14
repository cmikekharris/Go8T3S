package board

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cmikekharris/Go8T3S/internal/manual"
)

func Run() {
	// Make tree and starting board
	root := makeBranch([3][3]int{{-2}}, -1, &branch{})
	// Current branch
	cb := &root

	var choice string

	fmt.Println("Starting position :")

	printBoard(cb.Board)

	for {
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%s", &choice)

		switch choice {
		case "x":
			fmt.Println("Exiting...")
			return
		case "a":
			b := move(cb.Board, "left")
			cb = grow(b, cb)
		case "d":
			b := move(cb.Board, "right")
			cb = grow(b, cb)
		case "w":
			b := move(cb.Board, "up")
			cb = grow(b, cb)
		case "s":
			b := move(cb.Board, "down")
			cb = grow(b, cb)
		case "i":
			fmt.Println(manual.Instructions())
		case "p":
			printBoard(cb.Board)
		case "b":
			printBranch(*cb)
		}

		if checkWinningState(cb.Board) {
			fmt.Println("\t\t\t\t== WINNER WINNER WINNER ==")
			fmt.Println("Winning solution")
			fmt.Println("----------------")
			printTree(root)
			return
		}
	}
}

func makeBoard() [3][3]int {
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
	board := [3][3]int{}

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

func makeTestBoard() [3][3]int {
	return [3][3]int{
		{0, 2, 3},
		{1, 8, 4},
		{7, 6, 5},
	}
}

func goalState() [3][3]int {
	return [3][3]int{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	}
}

func checkWinningState(b [3][3]int) bool {
	return b == goalState()
}
