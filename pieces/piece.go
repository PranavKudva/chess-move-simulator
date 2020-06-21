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

