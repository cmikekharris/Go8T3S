package main

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/internal/game"
	"github.com/cmikekharris/Go8T3S/internal/manual"
	"github.com/cmikekharris/Go8T3S/internal/output"
	"github.com/cmikekharris/Go8T3S/internal/tree"
)

func main() {
	fmt.Println(manual.Banner())
	fmt.Println(manual.Instructions())

	var choice string
	var root = tree.Branch{}
	cb := root.Grow(game.MakeBoard())

	// Start game running loop
	for {
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%s", &choice)

		switch choice {
		case "x":
			fmt.Println("Exiting...")
			return
		case "a":
			b := cb.Board.Move("left")
			cb = cb.Grow(b)
		case "d":
			b := cb.Board.Move("right")
			cb = cb.Grow(b)
		case "w":
			b := cb.Board.Move("up")
			cb = cb.Grow(b)
		case "s":
			b := cb.Board.Move("down")
			cb = cb.Grow(b)
		case "i":
			fmt.Println(manual.Instructions())
		case "p":
			output.DisplayBoard(cb.Board)
		case "b":
			output.DisplayBranch(*cb)
		case "t":
			output.DisplayTree(root)
		}

		if game.WinningState(cb.Board) {
			fmt.Println("\t\t\t\t== WINNER WINNER WINNER ==")
			fmt.Println("Winning solution")
			fmt.Println("----------------")
			output.DisplayTree(root)
			return
		}
	}
}
