package test

import (
	"testing"

	"github.com/cmikekharris/Go8T3S/internal/game"
)

func TestMoveLeft(t *testing.T) {
	t.Log("Given the need to test moving the blank space left...")

	got := game.GameBoard{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}
	want := game.GameBoard{{1, 2, 3}, {0, 8, 4}, {7, 6, 5}}

	got = got.Move("left")

	if got != want {
		t.Fatal("\t\tMove the empty space left.", ballotX)
	}

	t.Log("\t\tMove the empty space left.", checkMark)
}

func TestMoveRight(t *testing.T) {
	t.Log("Given the need to test moving the blank space right...")

	got := game.GameBoard{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}
	want := game.GameBoard{{1, 2, 3}, {8, 4, 0}, {7, 6, 5}}

	got = got.Move("right")

	if got != want {
		t.Fatal("\t\tMove the empty space right.", ballotX)
	}

	t.Log("\t\tMove the empty space right.", checkMark)
}

func TestMoveUp(t *testing.T) {
	t.Log("Given the need to test moving the blank space up...")

	got := game.GameBoard{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}
	want := game.GameBoard{{1, 0, 3}, {8, 2, 4}, {7, 6, 5}}

	got = got.Move("up")

	if got != want {
		t.Fatal("\t\tMove the empty space up.", ballotX)
	}

	t.Log("\t\tMove the empty space up.", checkMark)
}

func TestMoveDown(t *testing.T) {
	t.Log("Given the need to test moving the blank space down...")

	got := game.GameBoard{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}
	want := game.GameBoard{{1, 2, 3}, {8, 6, 4}, {7, 0, 5}}

	got = got.Move("down")

	if got != want {
		t.Fatal("\t\tMove the empty space down.", ballotX)
	}

	t.Log("\t\tMove the empty space down.", checkMark)
}

func TestMoveOutOfBounds(t *testing.T) {
	t.Log("Given the need to test trying to move out of bounds of the board...")

	got := game.GameBoard{{1, 2, 3}, {8, 6, 4}, {7, 0, 5}}
	want := game.GameBoard{{1, 2, 3}, {8, 6, 4}, {7, 0, 5}}

	got = got.Move("down")

	if got != want {
		t.Fatal("\t\tDo no move out of bounds of the board.", ballotX)
	}

	t.Log("\t\tDo no move out of bounds of the board.", checkMark)
}
