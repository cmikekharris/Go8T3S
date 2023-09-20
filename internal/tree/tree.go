package tree

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/internal/game"
)

type Branch struct {
	Parent *Branch
	Moves  []*Branch
	Cost   int
	Board  game.GameBoard
}

// Store the new board in the tree and advance current pointer
func (cb *Branch) Grow(b [3][3]int) *Branch {
	nb := makeBranch(b, cb)
	cb.Moves = append(cb.Moves, &nb)

	return &nb
}

func makeBranch(board [3][3]int, parent *Branch) Branch {
	fmt.Println("Generating branch...")

	var branch = Branch{}
	branch.Board = board
	branch.Cost = parent.Cost + 1
	branch.Parent = parent

	return branch
}
