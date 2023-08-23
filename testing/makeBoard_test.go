package testing

import (
	"testing"

	"github.com/cmikekharris/Go8T3S/board"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestMakeBoard(t *testing.T) {
	t.Log("Given the need to test creating a new start board...")

	b := board.MakeBoard()
	un := [9]int{}
	count := 0

	t.Logf("\tWhen checking all created numbers...")

	for row, column := range b {
		for cell := range column {
			count++

			if b[row][cell] > 8 || b[row][cell] < 0 {
				t.Fatal("\t\tGenerated board numbers in correct bounds.", ballotX)
			}

			if un[b[row][cell]] == 1 {
				t.Fatal("\t\tGenerated board number unique.", ballotX)
			}

			un[b[row][cell]] = 1
		}
	}

	t.Log("\t\tGenerated board numbers in correct bounds.", checkMark)
	t.Log("\t\tGenerated board numbers unique.", checkMark)

	if count < 9 {
		t.Fatalf("\t\tCorrect number of numbers generated. Wanted 9, got %d %s", count, ballotX)
	}

	t.Log("\t\tCorrect number of numbers generated.", checkMark)
}
