package main

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/board"
	"github.com/cmikekharris/Go8T3S/tree"
)

func main() {
	fmt.Println(banner())
	fmt.Println(instructions())

	// Make a starting board
	root := tree.MakeBranch([3][3]int{{-2}}, -1, &tree.Tree{})
	// Current branch
	cb := &root

	tree.PrintBoard(cb.Board)

	b := [3][3]int{{-1}}
	var choice string

	for {
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%s", &choice)

		switch choice {
		case "x":
			fmt.Println("Exiting...")
			return
		case "a":
			b = board.Move(cb.Board, "left")
		case "d":
			b = board.Move(cb.Board, "right")
		case "w":
			b = board.Move(cb.Board, "up")
		case "s":
			b = board.Move(cb.Board, "down")
		case "i":
			fmt.Println(instructions())
		case "p":
			tree.PrintBoard(cb.Board)
		}

		// Store the new board in the tree and advance current pointer
		nb := tree.MakeBranch(b, cb.Cost, cb)
		cb.Moves = append(cb.Moves, &nb)
		cb = &nb

		tree.PrintBoard(cb.Board)

		if tree.CheckWinningState(*cb) {
			fmt.Println("\t\t\t\t== WINNER WINNER WINNER ==")
			fmt.Println("Winning solution")
			fmt.Println("----------------")
			tree.PrintTree(root)
			return
		}
	}
}

func banner() string {
	return `
	$$$$$$$$\ $$\           $$\        $$\           $$$$$$$$\ $$\ $$\                  $$$$$$\                                    
	$$  _____|\__|          $$ |       $$ |          \__$$  __|\__|$$ |                $$  __$$\                                   
	$$ |      $$\  $$$$$$\  $$$$$$$\ $$$$$$\            $$ |   $$\ $$ | $$$$$$\        $$ /  \__| $$$$$$\  $$$$$$\$$$$\   $$$$$$\  
	$$$$$\    $$ |$$  __$$\ $$  __$$\\_$$  _|           $$ |   $$ |$$ |$$  __$$\       $$ |$$$$\  \____$$\ $$  _$$  _$$\ $$  __$$\ 
	$$  __|   $$ |$$ /  $$ |$$ |  $$ | $$ |             $$ |   $$ |$$ |$$$$$$$$ |      $$ |\_$$ | $$$$$$$ |$$ / $$ / $$ |$$$$$$$$ |
	$$ |      $$ |$$ |  $$ |$$ |  $$ | $$ |$$\          $$ |   $$ |$$ |$$   ____|      $$ |  $$ |$$  __$$ |$$ | $$ | $$ |$$   ____|
	$$$$$$$$\ $$ |\$$$$$$$ |$$ |  $$ | \$$$$  |         $$ |   $$ |$$ |\$$$$$$$\       \$$$$$$  |\$$$$$$$ |$$ | $$ | $$ |\$$$$$$$\ 
	\________|\__| \____$$ |\__|  \__|  \____/          \__|   \__|\__| \_______|       \______/  \_______|\__| \__| \__| \_______|
                      $$\   $$ |                                                                                                       
                      \$$$$$$  |                                                                                                       
                       \______/	                                             
	`
}

func instructions() string {
	return `
	w - Up
	s - Down
	a - Left
	d - Right
	i - Print instructions
	x - Exit game
	`
}
