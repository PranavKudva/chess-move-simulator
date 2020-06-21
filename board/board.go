package board

type Grid struct {
	VerticalSize, HorizontalSize int
}

type Board Grid

func (board Board) Has(cell Cell) bool {
	return cell.HorizontalCoordinate >= 0 && cell.HorizontalCoordinate < board.HorizontalSize &&
		cell.VerticalCoordinate >= 0 && cell.VerticalCoordinate < board.VerticalSize
}

func BuildBoard(verticalSize, horizontalSize int) Board {
	return Board{verticalSize + 1, horizontalSize + 1}
}
