package traversals

import (
	"chess/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TraverseSingleMovements(t *testing.T) {
	tests := []struct {
		name        string
		direction   board.Direction
		acrossBoard bool
		currentCell board.Cell
		board       board.Board
		expected    board.Cells
	}{
		{
			"should return correct cell after performing the given movement",
			board.Up, false, board.Cell{2, 2}, board.Board{4, 4}, board.Cells{{1, 2}},
		},
		{
			"should return multiple cells when across board",
			board.Up, true, board.Cell{2, 2}, board.Board{4, 4}, board.Cells{{1, 2}, {0, 2}},
		},
		{
			"should return empty when movement not possible",
			board.Up, true, board.Cell{0, 2}, board.Board{4, 4}, board.Cells{},
		},
	}
	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			assert.Equal(t, subtest.expected,
				TraverseSingleMovements(subtest.direction, subtest.acrossBoard, subtest.currentCell, subtest.board))
		})
	}
}

func Test_TraverseMultiMovements(t *testing.T) {
	tests := []struct {
		name        string
		movement    Movement
		currentCell board.Cell
		board       board.Board
		expected    *board.Cell
	}{
		{
			"should return correct cell after performing the given movements",
			Movement{{board.Left, 2}, {board.Up, 1}}, board.Cell{2, 2}, board.Board{4, 4}, &board.Cell{1, 0},
		},
		{
			"should return nil when out of board",
			Movement{{board.Left, 2}, {board.Up, 1}}, board.Cell{0, 0}, board.Board{1, 1}, nil,
		},
		{
			"should return current when no movement",
			Movement{}, board.Cell{0, 0}, board.Board{1, 1}, &board.Cell{0, 0},
		},
		{
			"should return nil when complete movement not possible",
			Movement{{board.Left, 2}, {board.Up, 1}}, board.Cell{0, 2}, board.Board{4, 4}, nil,
		},
	}
	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			assert.Equal(t, subtest.expected, TraverseMultiMovements(subtest.movement, subtest.currentCell, subtest.board))
		})
	}
}
