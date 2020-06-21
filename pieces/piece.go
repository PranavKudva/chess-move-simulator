package piece

import (
	"chess/board"
	"chess/traversals"
)

type Piece struct {
	name               string
	board              board.Board
	current            board.Cell
	canMoveAcrossBoard bool
}

type HorizontalMovable interface {
	moveHorizontally() board.Cells
}

type HorizontalMovablePiece struct {
	Piece
}

func (piece HorizontalMovablePiece) moveHorizontally() board.Cells {
	nextCells := board.Cells{}
	for _, horizontalDirection := range board.HorizontalDirections {
		nextCells = append(nextCells,
			traversals.TraverseSingleMovements(horizontalDirection, piece.canMoveAcrossBoard, piece.current, piece.board)...)
	}

	return nextCells
}

type VerticalMovable interface {
	moveVertically() board.Cells
}

type VerticalMovablePiece struct {
	Piece
}

func (piece VerticalMovablePiece) moveVertically() board.Cells {
	nextCells := board.Cells{}
	for _, verticalDirection := range board.VerticalDirections {
		nextCells = append(nextCells,
			traversals.TraverseSingleMovements(verticalDirection, piece.canMoveAcrossBoard, piece.current, piece.board)...)
	}

	return nextCells
}

