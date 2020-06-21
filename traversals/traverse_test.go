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

