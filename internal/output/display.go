package output

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/internal/tree"
)

func DisplayBoard(b [3][3]int) {
	fmt.Printf("\t\tCurrent board\t\t\t\t\tWinning Board\n")
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\t\t\t\t\t| 1 | 2 | 3 |\n", b[0][0], b[0][1], b[0][2])
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\t\t\t\t\t| 8 | 0 | 4 |\n", b[1][0], b[1][1], b[1][2])
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
	fmt.Printf("\t\t| %d | %d | %d |\t\t\t\t\t| 7 | 6 | 5 |\n", b[2][0], b[2][1], b[2][2])
	fmt.Printf("\t\t-------------\t\t\t\t\t-------------\n")
}

func DisplayBranch(b tree.Branch) {
	fmt.Println("Printing branch...")
	fmt.Println("Parent : ", b.Parent)
	fmt.Println("Moves : ", b.Moves)
	fmt.Println("Cost : ", b.Cost)
	DisplayBoard(b.Board)
}

func DisplayTree(b tree.Branch) {
	fmt.Println("Printing tree...")
	fmt.Println("Number of branches : ", len(b.Moves))
	fmt.Printf("Parent : %p\n", b.Parent)

	for run := true; run; {
		fmt.Println("Cost : ", b.Cost)
		DisplayBoard(b.Board)

		for _, branch := range b.Moves {
			DisplayTree(*branch)
		}
		run = false
	}
}
