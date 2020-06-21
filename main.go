package main

import (
	"chess/board"
	"chess/pieces"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:3]
	pieceType, currCell := args[0], args[1]

	var nextCells []string
	eightByEightBoard := board.BuildBoard(8, 8)
	for _, cell := range getNextCells(eightByEightBoard, pieceType, currCell) {
		if cell.VerticalCoordinate > 0 {
			nextCells = append(nextCells, cell.String())
		}
	}

	fmt.Println(strings.Join(nextCells, ", "))
}

func strToCell(b board.Board, cellStr string) board.Cell {
	cellInUpper := strings.ToUpper(cellStr)
	rowColumnPair := strings.Split(cellInUpper, "")

	vertical := int(rowColumnPair[0][0]) % 65

	horizontal, err := strconv.Atoi(rowColumnPair[1])
	if err != nil {
		panic(fmt.Sprintf("horizontal coordinate %s is not a number", rowColumnPair[1]))
	}

	verticalCoordinate, horizontalCoordinate := vertical+1, horizontal-1

	cell := board.Cell{verticalCoordinate, horizontalCoordinate}
	if !b.Has(cell) {
		panic(fmt.Sprintf("8 X 8 b doesn't have the given coordinates"))
	}

	return cell
}

func getNextCells(b board.Board, pieceType string, currCell string) board.Cells {
	currentCell := strToCell(b, currCell)

	var pieceOnBoard piece.Movable

	switch strings.ToLower(pieceType) {
	case "king":
		pieceOnBoard = piece.King(b, currentCell)
	case "queen":
		pieceOnBoard = piece.Queen(b, currentCell)
	case "bishop":
		pieceOnBoard = piece.Bishop(b, currentCell)
	case "rook":
		pieceOnBoard = piece.Rook(b, currentCell)
	case "pawn":
		pieceOnBoard = piece.Pawn(b, currentCell)
	case "horse":
		pieceOnBoard = piece.Horse(b, currentCell)
	default:
		panic(errors.New(fmt.Sprintf("unknown piece type %s", pieceType)))
	}

	return pieceOnBoard.GetNextCells()
}
