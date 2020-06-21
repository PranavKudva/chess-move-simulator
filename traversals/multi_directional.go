package traversals

import (
	"chess/board"
)

type Step struct {
	direction board.Direction
	count     int
}

type Movement []Step
