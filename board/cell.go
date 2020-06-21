package board

import "fmt"

type Cell struct {
	VerticalCoordinate, HorizontalCoordinate int
}

func (c Cell) String() string {
	if c.VerticalCoordinate == 0 {
		return fmt.Sprintf("%c%d", 65 + c.VerticalCoordinate, c.HorizontalCoordinate + 1)
	}

	return fmt.Sprintf("%c%d", 65 + (c.VerticalCoordinate - 1), c.HorizontalCoordinate + 1)
}

func (c Cell) leftCell(count int) Cell {
	return Cell{HorizontalCoordinate: c.HorizontalCoordinate - count, VerticalCoordinate: c.VerticalCoordinate}
}

func (c Cell) rightCell(count int) Cell {
	return Cell{HorizontalCoordinate: c.HorizontalCoordinate + count, VerticalCoordinate: c.VerticalCoordinate}
}

func (c Cell) upCell(count int) Cell {
	return Cell{HorizontalCoordinate: c.HorizontalCoordinate, VerticalCoordinate: c.VerticalCoordinate - count}
}

func (c Cell) downCell(count int) Cell {
	return Cell{HorizontalCoordinate: c.HorizontalCoordinate, VerticalCoordinate: c.VerticalCoordinate + count}
}

func (c Cell) leftUpCell(count int) Cell {
	return Cell{HorizontalCoordinate: c.HorizontalCoordinate - count, VerticalCoordinate: c.VerticalCoordinate - count}
}

func (c Cell) rightUpCell(count int) Cell {
	return Cell{HorizontalCoordinate: c.HorizontalCoordinate + count, VerticalCoordinate: c.VerticalCoordinate - count}
}

func (c Cell) leftDownCell(count int) Cell {
	return Cell{HorizontalCoordinate: c.HorizontalCoordinate - count, VerticalCoordinate: c.VerticalCoordinate + count}
}

func (c Cell) rightDownCell(count int) Cell {
	return Cell{HorizontalCoordinate: c.HorizontalCoordinate + count, VerticalCoordinate: c.VerticalCoordinate + count}
}

type Cells []Cell
