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

func TestPiece_moveVertically(t *testing.T) {
	tests := []struct {
		name     string
		piece    VerticalMovablePiece
		expected board.Cells
	}{
		// A1
		// B1
		// C1
		// D1
		// E1; current = C1
		{
			name:     "piece on 5X1 board should return apt possible vertical cells when can't move across board",
			piece:    VerticalMovablePiece{Piece{board: board.Board{5, 1}, current: board.Cell{2, 0}, canMoveAcrossBoard: false}},
			expected: board.Cells{board.Cell{1, 0}, board.Cell{3, 0}},
		},
		{
			name:     "piece on 5X1 board should return all cells to vertical corner when can move across board",
			piece:    VerticalMovablePiece{Piece{board: board.Board{5, 1}, current: board.Cell{2, 0}, canMoveAcrossBoard: true}},
			expected: board.Cells{board.Cell{1, 0}, board.Cell{0, 0}, board.Cell{3, 0}, board.Cell{4, 0}},
		},
		// A1
		// B1
		// C1
		// D1
		// E1; current = A1
		{
			name:     "piece on 5X1 board should return apt possible vertical cells when can't move up",
			piece:    VerticalMovablePiece{Piece{board: board.Board{5, 1}, current: board.Cell{0, 0}, canMoveAcrossBoard: true}},
			expected: board.Cells{{1, 0}, {2, 0}, {3, 0}, {4, 0}},
		},
		// A1
		// B1
		// C1
		// D1
		// E1; current = E1
		{
			name:     "piece on 5X1 board should return apt possible vertical cells when can't move down",
			piece:    VerticalMovablePiece{Piece{board: board.Board{5, 1}, current: board.Cell{4, 0}, canMoveAcrossBoard: true}},
			expected: board.Cells{{3, 0}, {2, 0}, {1, 0}, {0, 0}},
		},
		// A1; current = A1
		{
			name:     "piece on 1X1 board should return empty when on vertical corner and can't move across board",
			piece:    VerticalMovablePiece{Piece{board: board.Board{1, 1}, current: board.Cell{0, 0}, canMoveAcrossBoard: false}},
			expected: board.Cells{},
		},
		{
			name:     "piece on 1X1 board should return empty when on vertical corner and can move across board",
			piece:    VerticalMovablePiece{Piece{board: board.Board{1, 1}, current: board.Cell{0, 0}, canMoveAcrossBoard: true}},
			expected: board.Cells{},
		},
	}
	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			cells := subtest.piece.moveVertically()
			assert.Equal(t, subtest.expected, cells)
		})
	}
}

func TestPiece_moveDiagonally(t *testing.T) {
	tests := []struct {
		name     string
		piece    DiagonalMovablePiece
		expected board.Cells
	}{
		// A1 A2 A3 A4 A5
		// B1 B2 B3 B4 B5
		// C1 C2 C3 C4 C5
		// D1 D2 D3 D4 D5
		// E1 E2 E3 E4 E5; current = C3
		{
			name:     "piece on 5X5 board should return apt possible diagonal cells when can't move across board",
			piece:    DiagonalMovablePiece{Piece{board: board.Board{5, 5}, current: board.Cell{2, 2}, canMoveAcrossBoard: false}},
			expected: board.Cells{board.Cell{1, 1}, board.Cell{3, 1}, board.Cell{1, 3}, board.Cell{3, 3}},
		},
		{
			name:     "piece on 5X5 board should return apt possible diagonal cells when can move across board",
			piece:    DiagonalMovablePiece{Piece{board: board.Board{5, 5}, current: board.Cell{2, 2}, canMoveAcrossBoard: true}},
			expected: board.Cells{board.Cell{1, 1}, board.Cell{0, 0}, board.Cell{3, 1}, board.Cell{4, 0}, board.Cell{1, 3}, board.Cell{0, 4}, board.Cell{3, 3}, board.Cell{4, 4}},
		},
		// A1 A2 A3 A4 A5
		// B1 B2 B3 B4 B5
		// C1 C2 C3 C4 C5
		// D1 D2 D3 D4 D5
		// E1 E2 E3 E4 E5; current = A1
		{
			name:     "piece on 5X5 board should return apt possible diagonal cells when can't move up left",
			piece:    DiagonalMovablePiece{Piece{board: board.Board{5, 5}, current: board.Cell{0, 0}, canMoveAcrossBoard: true}},
			expected: board.Cells{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
		},
		// A1 A2 A3 A4 A5
		// B1 B2 B3 B4 B5
		// C1 C2 C3 C4 C5
		// D1 D2 D3 D4 D5
		// E1 E2 E3 E4 E5; current = E5
		{
			name:     "piece on 5X5 board should return apt possible diagonal cells when can't move down right",
			piece:    DiagonalMovablePiece{Piece{board: board.Board{5, 5}, current: board.Cell{4, 4}, canMoveAcrossBoard: true}},
			expected: board.Cells{{3, 3}, {2, 2}, {1, 1}, {0, 0}},
		},
		// A1; current = A1
		{
			name:     "piece on 1X1 board should return empty when on diagonal corner and can't move across board",
			piece:    DiagonalMovablePiece{Piece{board: board.Board{1, 1}, current: board.Cell{0, 0}, canMoveAcrossBoard: false}},
			expected: board.Cells{},
		},
		{
			name:     "piece on 1X1 board should return empty when on diagonal corner and can move across board",
			piece:    DiagonalMovablePiece{Piece{board: board.Board{1, 1}, current: board.Cell{0, 0}, canMoveAcrossBoard: true}},
			expected: board.Cells{},
		},
	}
	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			cells := subtest.piece.moveDiagonally()
			assert.Equal(t, subtest.expected, cells)
		})
	}
}
