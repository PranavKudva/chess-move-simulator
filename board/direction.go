package board

type Direction func(cell Cell, stepCount int) Cell

var Left = func(cell Cell, stepCount int) Cell {
	return cell.leftCell(stepCount)
}

var Right = func(cell Cell, stepCount int) Cell {
	return cell.rightCell(stepCount)
}

var Up = func(cell Cell, stepCount int) Cell {
	return cell.upCell(stepCount)
}

var Down = func(cell Cell, stepCount int) Cell {
	return cell.downCell(stepCount)
}

var LeftUp = func(cell Cell, stepCount int) Cell {
	return cell.leftUpCell(stepCount)
}

var RightUp = func(cell Cell, stepCount int) Cell {
	return cell.rightUpCell(stepCount)
}

var LeftDown = func(cell Cell, stepCount int) Cell {
	return cell.leftDownCell(stepCount)
}

var RightDown = func(cell Cell, stepCount int) Cell {
	return cell.rightDownCell(stepCount)
}

var HorizontalDirections = []Direction{ Left, Right }

var VerticalDirections = []Direction{ Up, Down }

