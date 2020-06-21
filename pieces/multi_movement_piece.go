package piece

import (
	"chess/board"
	"chess/traversals"
)

type MultiStepHorizontalMovablePiece struct {
	Piece
}

func (piece MultiStepHorizontalMovablePiece) moveHorizontally() board.Cells {
	nextCells := board.Cells{}

	for _, moves := range traversals.GetPossibleMovements(2, 1, board.HorizontalDirections, board.VerticalDirections) {
		if cell := traversals.TraverseMultiMovements(moves, piece.current, piece.board); cell != nil {
			nextCells = append(nextCells, *cell)
		}
	}

	return nextCells
}

type MultiStepVerticalMovablePiece struct {
	Piece
}

func (piece MultiStepVerticalMovablePiece) moveVertically() board.Cells {
	nextCells := board.Cells{}

	for _, moves := range traversals.GetPossibleMovements(2, 1, board.VerticalDirections, board.HorizontalDirections) {
		if cell := traversals.TraverseMultiMovements(moves, piece.current, piece.board); cell != nil {
			nextCells = append(nextCells, *cell)
		}
	}

	return nextCells
}

