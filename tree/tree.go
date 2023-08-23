package tree

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/board"
)

type Tree struct {
	move1 *Tree
	move2 *Tree
	move3 *Tree
	move4 *Tree
	board [3][3]int
}

func GenerateTree() *Tree {
	fmt.Println("Generating Tree...")

	t := Tree{}

	t.board = board.GenerateNewBoard()

	return &t
}
