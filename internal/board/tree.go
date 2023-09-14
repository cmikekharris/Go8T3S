package board

import (
	"fmt"
)

type branch struct {
	Parent *branch
	Moves  []*branch
	Cost   int
	Board  [3][3]int
}

func makeBranch(board [3][3]int, cost int, parent *branch) branch {
	fmt.Println("Generating branch...")

	var branch = branch{}
	branch.Board = board
	branch.Cost = cost + 1
	branch.Parent = parent

	if branch.Board[0][0] == -1 {
		branch.Board = makeBoard()
	}

	if branch.Board[0][0] == -2 {
		branch.Board = makeTestBoard()
	}

	return branch
}

func grow(b [3][3]int, cb *branch) *branch {
	// Store the new board in the tree and advance current pointer
	nb := makeBranch(b, cb.Cost, cb)
	cb.Moves = append(cb.Moves, &nb)
	return &nb
}

func printTree(b branch) {
	fmt.Println("Printing tree...")
	fmt.Println("Number of branches : ", len(b.Moves))
	fmt.Printf("Parent : %p\n", b.Parent)

	for run := true; run; {
		fmt.Println("Cost : ", b.Cost)
		printBoard(b.Board)

		for _, branch := range b.Moves {
			printTree(*branch)
		}
		run = false
	}
}

func printBranch(b branch) {
	fmt.Println("Printing branch...")
	fmt.Println("Parent : ", b.Parent)
	fmt.Println("Moves : ", b.Moves)
	fmt.Println("Cost : ", b.Cost)
	printBoard(b.Board)
}

func printBoard(b [3][3]int) {
	fmt.Printf("\t\tCurrent board\t\t\t\t\tWinning Board\n")
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\t\t\t\t\t| 1 | 2 | 3 |\n", b[0][0], b[0][1], b[0][2])
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\t\t\t\t\t| 8 | 0 | 4 |\n", b[1][0], b[1][1], b[1][2])
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\t\t\t\t\t| 7 | 6 | 5 |\n", b[2][0], b[2][1], b[2][2])
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
}
