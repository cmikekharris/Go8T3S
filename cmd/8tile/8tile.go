package main

import (
	"fmt"

	"github.com/cmikekharris/Go8T3S/internal/board"
	"github.com/cmikekharris/Go8T3S/internal/manual"
)

func main() {
	fmt.Println(manual.Banner())
	fmt.Println(manual.Instructions())
	
	board.Run()
}
