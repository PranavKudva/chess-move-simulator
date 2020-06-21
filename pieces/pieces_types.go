package piece

import "chess/board"

func Pawn(board board.Board, cell board.Cell) VerticalMovablePiece {
	return VerticalMovablePiece{Piece{"Pawn", board, cell, false}}
}

func Bishop(board board.Board, cell board.Cell) DiagonalMovablePiece {
	piece := Piece{"Bishop", board, cell, true}

	return DiagonalMovablePiece{piece}
}

func King(board board.Board, cell board.Cell) AllMovablePiece {
	piece := Piece{"King", board, cell, false}

	return AllMovablePiece{
		DiagonalMovablePiece{piece},
		VerticalMovablePiece{piece},
		HorizontalMovablePiece{piece},
	}
}

func Queen(board board.Board, cell board.Cell) AllMovablePiece {
	piece := Piece{"Queen", board, cell, true}

	return AllMovablePiece{
		DiagonalMovablePiece{piece},
		VerticalMovablePiece{piece},
		HorizontalMovablePiece{piece},
	}
}

func Rook(board board.Board, cell board.Cell) HorizontalAndVerticalMovablePiece {
	piece := Piece{"Bishop", board, cell, true}

	return HorizontalAndVerticalMovablePiece{
		VerticalMovablePiece{piece},
		HorizontalMovablePiece{piece},
	}
}

func Horse(board board.Board, cell board.Cell) MultiMovementPiece {
	piece := Piece{"Knight", board, cell, true}

	return MultiMovementPiece{
		MultiStepHorizontalMovablePiece{piece},
		MultiStepVerticalMovablePiece{piece},
	}
}
