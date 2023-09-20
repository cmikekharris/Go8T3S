package test

import (
	"testing"

	"github.com/cmikekharris/Go8T3S/internal/game"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestMakeBoard(t *testing.T) {
	t.Log("Given the need to test creating a new random start board...")

	b := game.MakeBoard()
	un := [9]int{}
	count := 0

	t.Log("\tWhen checking all created numbers...")

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

func TestMakeTestBoard(t *testing.T) {
	t.Log("Given the need to test creating a test start board...")

	b := game.MakeTestBoard()
	want := [3][3]int{
		{0, 2, 3},
		{1, 8, 4},
		{7, 6, 5},
	}

	if b != want {
		t.Fatalf("\t\tPredictable test start board generated. %s", ballotX)
	}

	t.Log("\t\tPredictable test start board generated.", checkMark)
}

func TestWinningState(t *testing.T) {
	t.Log("Given the need to check for the winning state...")

	b := [3][3]int{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	}

	if !game.WinningState(b) {
		t.Fatalf("\t\tWinning state corrently set. %s", ballotX)
	}

	t.Log("\t\tWinning state corrently set.", checkMark)
}
