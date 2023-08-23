package main

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/tree"
)

func main() {
	fmt.Println("Eight Tile State Space Searching Game written in Go - Go8T3S")

	root := tree.MakeTree()

	fmt.Println(root)
}
