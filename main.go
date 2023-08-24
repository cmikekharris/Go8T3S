package main

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/board"
	"github.com/cmikekharris/Go8T3S/tree"
)

func main() {
	fmt.Println("Eight Tile State Space Searching Game written in Go - Go8T3S")

	root := tree.MakeTree()

	fmt.Println(root)
	tree.PrintTree(root)

	root.Board = board.Move(root.Board, "up")
	tree.PrintTree(root)

	root.Board = board.Move(root.Board, "down")
	tree.PrintTree(root)

	root.Board = board.Move(root.Board, "left")
	tree.PrintTree(root)

	root.Board = board.Move(root.Board, "right")
	tree.PrintTree(root)
}
