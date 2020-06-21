package piece

import (
	"chess/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPiece_moveHorizontally(t *testing.T) {
	tests := []struct {
		name     string
		piece    HorizontalMovablePiece
		expected board.Cells
	}{
		// A1 A2 A3 A4 A5; current = A3
		{
			name:     "piece on 1X5 board should return apt possible horizontal cells when can't move across board",
			piece:    HorizontalMovablePiece{Piece{board: board.Board{1, 5}, canMoveAcrossBoard: false, current: board.Cell{0, 2}}},
			expected: board.Cells{board.Cell{0, 1}, board.Cell{0, 3}},
		},
		{
			name:     "piece on 1X5 board should return all cells to horizontal corner when can move across board",
			piece:    HorizontalMovablePiece{Piece{board: board.Board{1, 5}, current: board.Cell{0, 2}, canMoveAcrossBoard: true}},
			expected: board.Cells{{0, 1}, {0, 0}, {0, 3}, {0, 4}},
		},
		// A1 A2 A3 A4 A5; current = A1,
		{
			name:     "piece on 1X5 board should return apt possible horizontal cells when can't move left",
			piece:    HorizontalMovablePiece{Piece{board: board.Board{1, 5}, current: board.Cell{0, 0}, canMoveAcrossBoard: true}},
			expected: board.Cells{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
		},
		// A1 A2 A3 A4 A5; current = A5,
		{
			name:     "piece on 1X5 board should return apt possible horizontal cells when can't move right",
			piece:    HorizontalMovablePiece{Piece{board: board.Board{1, 5}, current: board.Cell{0, 4}, canMoveAcrossBoard: true}},
			expected: board.Cells{{0, 3}, {0, 2}, {0, 1}, {0, 0}},
		},
		// A1; current = A1
		{
			name:     "piece on 1X1 board should return empty when on horizontal corner and can't move across board",
			piece:    HorizontalMovablePiece{Piece{board: board.Board{1, 1}, current: board.Cell{0, 0}, canMoveAcrossBoard: false}},
			expected: board.Cells{},
		},
		// A1; current = A1
		{
			name:     "piece on 1X1 board should return empty when on horizontal corner and can move across board",
			piece:    HorizontalMovablePiece{Piece{board: board.Board{1, 1}, current: board.Cell{0, 0}, canMoveAcrossBoard: true}},
			expected: board.Cells{},
		},
	}
	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			cells := subtest.piece.moveHorizontally()
			assert.Equal(t, subtest.expected, cells)
		})
	}
}

