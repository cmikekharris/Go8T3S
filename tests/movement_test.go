package test

import (
	"testing"

	"github.com/cmikekharris/Go8T3S/internal/game"
)

func TestMove(t *testing.T) {
	t.Log("Given the need to test moving around the board...")

	var tests = []struct {
		name      string
		direction string
		input     game.GameBoard
		want      game.GameBoard
	}{
		{"Move up 1", "up", game.GameBoard{{1, 2, 3}, {0, 8, 4}, {7, 6, 5}}, game.GameBoard{{0, 2, 3}, {1, 8, 4}, {7, 6, 5}}},
		{"Move up 2", "up", game.GameBoard{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, game.GameBoard{{1, 0, 3}, {8, 2, 4}, {7, 6, 5}}},
		{"Move up 3", "up", game.GameBoard{{1, 2, 3}, {8, 4, 0}, {7, 6, 5}}, game.GameBoard{{1, 2, 0}, {8, 4, 3}, {7, 6, 5}}},
		{"Move down 1", "down", game.GameBoard{{1, 2, 3}, {0, 8, 4}, {7, 6, 5}}, game.GameBoard{{1, 2, 3}, {7, 8, 4}, {0, 6, 5}}},
		{"Move down 2", "down", game.GameBoard{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, game.GameBoard{{1, 2, 3}, {8, 6, 4}, {7, 0, 5}}},
		{"Move down 3", "down", game.GameBoard{{1, 2, 3}, {8, 4, 0}, {7, 6, 5}}, game.GameBoard{{1, 2, 3}, {8, 4, 5}, {7, 6, 0}}},
		{"Move left 1", "left", game.GameBoard{{1, 0, 3}, {8, 2, 4}, {7, 6, 5}}, game.GameBoard{{0, 1, 3}, {8, 2, 4}, {7, 6, 5}}},
		{"Move left 2", "left", game.GameBoard{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, game.GameBoard{{1, 2, 3}, {0, 8, 4}, {7, 6, 5}}},
		{"Move left 3", "left", game.GameBoard{{1, 2, 3}, {8, 6, 4}, {7, 0, 5}}, game.GameBoard{{1, 2, 3}, {8, 6, 4}, {0, 7, 5}}},
		{"Move right 1", "right", game.GameBoard{{1, 0, 3}, {8, 2, 4}, {7, 6, 5}}, game.GameBoard{{1, 3, 0}, {8, 2, 4}, {7, 6, 5}}},
		{"Move right 2", "right", game.GameBoard{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, game.GameBoard{{1, 2, 3}, {8, 4, 0}, {7, 6, 5}}},
		{"Move right 3", "right", game.GameBoard{{1, 2, 3}, {8, 6, 4}, {7, 0, 5}}, game.GameBoard{{1, 2, 3}, {8, 6, 4}, {7, 5, 0}}},
	}

	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Move(tt.direction)

			if got != tt.want {
				t.Fatalf("\t\tMove the empty space %s. %s", tt.direction, ballotX)
			}

			t.Logf("\t\tMove the empty space %s. %s", tt.direction, checkMark)
		})
	}
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
