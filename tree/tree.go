package tree

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/board"
)

type Tree struct {
	parent *Tree
	uo    *Tree
	down  *Tree
	left  *Tree
	right *Tree
	board [3][3]int
}

func MakeTree() *Tree {
	fmt.Println("Generating Tree...")

	t := Tree{}

	t.board = board.MakeBoard()

	return &t
}

func PrintTree() {

}
