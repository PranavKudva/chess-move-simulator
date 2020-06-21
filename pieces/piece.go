package piece

import (
	"chess/board"
	"chess/traversals"
)

type Movable interface {
	GetNextCells() board.Cells
}

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

func (piece VerticalMovablePiece) GetNextCells() board.Cells  {
	return piece.moveVertically()
}

type DiagonalMovable interface {
	moveDiagonally() board.Cells
}

type DiagonalMovablePiece struct {
	Piece
}

func (piece DiagonalMovablePiece) moveDiagonally() board.Cells {
	nextCells := board.Cells{}

	for _, diagonalDirection := range board.DiagonalDirections {
		nextCells = append(nextCells,
			traversals.TraverseSingleMovements(diagonalDirection, piece.canMoveAcrossBoard, piece.current, piece.board)...)
	}

	return nextCells
}

func (piece DiagonalMovablePiece) GetNextCells() board.Cells  {
	return piece.moveDiagonally()
}

type AllMovablePiece struct {
	DiagonalMovablePiece
	VerticalMovablePiece
	HorizontalMovablePiece
}

func (piece AllMovablePiece) GetNextCells() board.Cells  {
	allPossibleCells := board.Cells{}

	horizonalCells := piece.HorizontalMovablePiece.moveHorizontally()
	verticalCells := piece.VerticalMovablePiece.moveVertically()
	diagonalCells := piece.DiagonalMovablePiece.moveDiagonally()

	allPossibleCells = append(allPossibleCells, horizonalCells...)
	allPossibleCells = append(allPossibleCells, verticalCells...)
	allPossibleCells = append(allPossibleCells, diagonalCells...)

	return allPossibleCells
}

type HorizontalAndVerticalMovablePiece struct {
	VerticalMovablePiece
	HorizontalMovablePiece
}

func (piece HorizontalAndVerticalMovablePiece) GetNextCells() board.Cells  {
	return append(piece.HorizontalMovablePiece.moveHorizontally(), piece.VerticalMovablePiece.moveVertically()...)
}
