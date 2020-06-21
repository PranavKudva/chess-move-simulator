package traversals

import (
	"chess/board"
)

func getCellAt(current board.Cell, direction board.Direction, stepCount int) board.Cell {
	return direction(current, stepCount)
}

func TraverseSingleMovements(direction board.Direction, acrossBoard bool,
	current board.Cell, brd board.Board) board.Cells {

	nextCells := board.Cells{}

	for nextCell := getCellAt(current, direction, 1); brd.Has(nextCell); nextCell = getCellAt(nextCell, direction, 1) {
		nextCells = append(nextCells, nextCell)
		if !acrossBoard {
			break
		}
	}

	return nextCells
}
