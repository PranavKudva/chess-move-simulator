package traversals

import (
	"chess/board"
)

type Step struct {
	direction board.Direction
	count     int
}

type Movement []Step

func GetPossibleMovements(majorCount, minorCount int, majorDirections, minorDirections []board.Direction) []Movement {
	var movements []Movement

	for _, majorDirection := range majorDirections {
		for _, minorDirection := range minorDirections {
			movements = append(movements,
				Movement{
					{count: majorCount, direction: majorDirection}, {count: minorCount, direction: minorDirection},
				})
		}
	}

	return movements
}
