package piece

import "chess/board"

func Pawn(board board.Board, cell board.Cell) VerticalMovablePiece {
	return VerticalMovablePiece{Piece{"Pawn", board, cell, false}}
}
