package tree

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/board"
)

type Tree struct {
	Parent *Tree
	Up     *Tree
	Down   *Tree
	Left   *Tree
	Right  *Tree
	Cost   int
	Board  [3][3]int
}

func MakeTree() Tree {
	fmt.Println("Generating Tree...")

	t := Tree{}

	t.Board = board.MakeBoard()

	return t
}

func PrintTree(t Tree) {
	fmt.Printf("\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\n", t.Board[0][0], t.Board[0][1], t.Board[0][2])
	fmt.Printf("\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\n", t.Board[1][0], t.Board[1][1], t.Board[1][2])
	fmt.Printf("\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\n", t.Board[2][0], t.Board[2][1], t.Board[2][2])
	fmt.Printf("\t\t-------------\n")
}
