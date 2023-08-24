package testing

import (
	"testing"

	"github.com/cmikekharris/Go8T3S/board"
)

func TestMoveLeft(t *testing.T) {
	t.Log("Given the need to test moving the blank space left...")

	got := [3][3]int{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}
	want := [3][3]int{{1, 2, 3}, {0, 8, 4}, {7, 6, 5}}

	got = board.Move(got, "left")

	if got != want {
		t.Fatal("\t\tMove the empty space left.", ballotX)
	}

	t.Log("\t\tMove the empty space left.", checkMark)
}

func TestMoveRight(t *testing.T) {
	t.Log("Given the need to test moving the blank space right...")

	got := [3][3]int{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}
	want := [3][3]int{{1, 2, 3}, {8, 4, 0}, {7, 6, 5}}

	got = board.Move(got, "right")

	if got != want {
		t.Fatal("\t\tMove the empty space right.", ballotX)
	}

	t.Log("\t\tMove the empty space right.", checkMark)
}

func TestMoveUp(t *testing.T) {
	t.Log("Given the need to test moving the blank space up...")

	got := [3][3]int{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}
	want := [3][3]int{{1, 0, 3}, {8, 2, 4}, {7, 6, 5}}

	got = board.Move(got, "up")

	if got != want {
		t.Fatal("\t\tMove the empty space up.", ballotX)
	}

	t.Log("\t\tMove the empty space up.", checkMark)
}

func TestMoveDown(t *testing.T) {
	t.Log("Given the need to test moving the blank space down...")

	got := [3][3]int{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}
	want := [3][3]int{{1, 2, 3}, {8, 6, 4}, {7, 0, 5}}

	got = board.Move(got, "down")

	if got != want {
		t.Fatal("\t\tMove the empty space down.", ballotX)
	}

	t.Log("\t\tMove the empty space down.", checkMark)
}

func TestFindBlankSpace(t *testing.T) {
	t.Log("Given the need to test finding the empty space...")

	b := [3][3]int{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}

	gotRow, gotCell := board.FindEmptySpace(b)

	if gotRow != 1 || gotCell != 1 {
		t.Fatal("\t\tFind the empty space.", ballotX)
	}

	t.Log("\t\tFind the empty space.", checkMark)
}

func TestMoveOutOfBounds(t *testing.T) {
	t.Log("Given the need to test trying to move out of bounds of the board...")

	got := [3][3]int{{1, 2, 3}, {8, 6, 4}, {7, 0, 5}}
	want := [3][3]int{{1, 2, 3}, {8, 6, 4}, {7, 0, 5}}

	got = board.Move(got, "down")

	if got != want {
		t.Fatal("\t\tDo no move out of bounds of the board.", ballotX)
	}

	t.Log("\t\tDo no move out of bounds of the board.", checkMark)
}
