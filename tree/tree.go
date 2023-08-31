package tree

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/board"
)

type Tree struct {
	Parent *Tree
	Moves  []*Tree
	Cost   int
	Board  [3][3]int
}

func MakeBranch(b [3][3]int, cost int, parent *Tree) Tree {
	fmt.Println("Generating branch...")

	var t = Tree{}
	t.Board = b
	t.Cost = cost + 1
	t.Parent = parent

	if t.Board[0][0] == -1 {
		t.Board = board.MakeBoard()
	}

	if t.Board[0][0] == -2 {
		t.Board = board.MakeTestBoard()
	}

	return t
}

func PrintTree(t Tree) {
	fmt.Println("Printing tree...")
	fmt.Println("Number of branches : ", len(t.Moves))
	fmt.Printf("Parent : %p\n", t.Parent)

	for run := true; run; {
		fmt.Println("Cost : ", t.Cost)
		PrintBoard(t.Board)

		for _, branch := range t.Moves {
			PrintTree(*branch)
		}
		run = false
	}
}

func PrintBranch(t Tree) {
	fmt.Println("Printing branch...")
	fmt.Println("Parent : ", t.Parent)
	fmt.Println("Moves : ", t.Moves)
	fmt.Println("Cost : ", t.Cost)
	PrintBoard(t.Board)
}

func PrintBoard(b [3][3]int) {
	fmt.Printf("\t\tCurrent board\t\t\t\t\tWinning Board\n")
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\t\t\t\t\t| 1 | 2 | 3 |\n", b[0][0], b[0][1], b[0][2])
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\t\t\t\t\t| 8 | 0 | 4 |\n", b[1][0], b[1][1], b[1][2])
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\t\t\t\t\t| 7 | 6 | 5 |\n", b[2][0], b[2][1], b[2][2])
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
}

func CheckWinningState(t Tree) bool {
	return t.Board == board.GoalState()
}
