package piece

import (
	"chess/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiMovementPiece_moveHorizontally(t *testing.T) {
	tests := []struct {
		name     string
		piece    MultiStepHorizontalMovablePiece
		expected board.Cells
	}{
		// A1 A2 A3 A4 A5
		// B1 B2 B3 B4 B5
		// C1 C2 C3 C4 C5; current = B3
		{
			name:     "piece on 3X5 board should return apt possible cells",
			piece:    MultiStepHorizontalMovablePiece{
				Piece{board: board.Board{3, 5}, canMoveAcrossBoard: false, current: board.Cell{1, 2}},
			},
			expected: board.Cells{board.Cell{0, 0}, board.Cell{2, 0}, board.Cell{0, 4}, board.Cell{2, 4}},
		},
		// A1 A2 A3 A4 A5
		// B1 B2 B3 B4 B5
		// C1 C2 C3 C4 C5; current = B5
		{
			name:     "piece on 3X5 board should return apt possible cells when can't move right",
			piece:    MultiStepHorizontalMovablePiece{
				Piece{board: board.Board{5, 3}, canMoveAcrossBoard: false, current: board.Cell{1, 4}},
			},
			expected: board.Cells{board.Cell{0, 2}, board.Cell{2, 2}},
		},
		// current = B1
		{
			name:     "piece on 3X5 board should return apt possible cells when can't move left",
			piece:    MultiStepHorizontalMovablePiece{
				Piece{board: board.Board{3, 5}, canMoveAcrossBoard: false, current: board.Cell{1, 1}},
			},
			expected: board.Cells{board.Cell{0, 3}, board.Cell{2, 3}},
		},
		// A1; current = A1
		{
				//	name:     "piece on 1X1 board should return empty when on corner and can't move across board",
				//	piece:    MultiStepHorizontalMovablePiece{Piece{board: Board{1, 1}, current: Cell{0, 0}, canMoveAcrossBoard: false}},
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

