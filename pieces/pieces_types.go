package piece

import "chess/board"

func Pawn(board board.Board, cell board.Cell) VerticalMovablePiece {
	return VerticalMovablePiece{Piece{"Pawn", board, cell, false}}
}

func Bishop(board board.Board, cell board.Cell) DiagonalMovablePiece {
	piece := Piece{"Bishop", board, cell, true}

	return DiagonalMovablePiece{piece}
}
