package game

import (
	"fmt"

	"mo.example/utils"
)

type CellType int

const (
	EMPTY CellType = iota
	SAND
	WATER
	WALL
)

func (ct CellType) String() string {
	switch ct {
	case EMPTY:
		return "EMPTY"
	case SAND:
		return "SAND"
	case WATER:
		return "WATER"
	case WALL:
		return "WALL"
	}
	return "UNKNOWN"
}

type Grid struct {
	CellSize int
	Cols     int
	Rows     int
	Data     [][]CellType
	moved    [][]bool
}

func (gird *Grid) Init() {
	for i := 0; i < gird.Rows; i++ {
		gird.Data = append(gird.Data, make([]CellType, gird.Cols))
		gird.moved = append(gird.moved, make([]bool, gird.Cols))

	}
}
func (grid *Grid) Reset() {
	grid.moved = make([][]bool, 0)
	grid.Data = make([][]CellType, 0)
	grid.Init()
}
func (grid *Grid) IndexAt(x int, y int) (_row int, _col int, e error) {
	if x < 0 || y < 0 {
		return -1, -1, fmt.Errorf("out of bounds")
	}
	col := x / grid.CellSize
	row := y / grid.CellSize
	if col >= grid.Cols || row >= grid.Rows {
		return -1, -1, fmt.Errorf("out of bounds")
	}
	return row, col, nil
}

func (grid *Grid) Step() {
	for i := range grid.moved {
		clear(grid.moved[i])
	}
	for i := grid.Rows - 2; i >= 0; i-- {
		for j := 0; j < grid.Cols; j++ {
			if grid.moved[i][j] {
				continue
			}
			switch grid.Data[i][j] {
			case EMPTY:
			case SAND:
				grid.updateSand(i, j)
			case WATER:
				grid.updateWater(i, j)
			}
		}
	}
}
func (grid *Grid) updateSand(row int, col int) {
	leftOrRight := 1
	if utils.RandomBool() {
		leftOrRight = -1
	}
	if grid.canSandEnter(row+1, col) {
		grid.swap(row, col, row+1, col)
		return
	}
	if grid.canSandEnter(row+1, col+leftOrRight) {
		grid.swap(row, col, row+1, col+leftOrRight)
		return
	}
	if grid.canSandEnter(row+1, col-leftOrRight) {
		grid.swap(row, col, row+1, col-leftOrRight)
	}
}
func (grid *Grid) updateWater(row int, col int) {
	leftOrRight := 1
	if utils.RandomBool() {
		leftOrRight = -1
	}
	if grid.canWaterEnter(row+1, col) {
		grid.swap(row, col, row+1, col)
		return
	}
	if grid.canSandEnter(row, col+leftOrRight) {
		grid.swap(row, col, row, col+leftOrRight)
	}
}

func (grid *Grid) At(row int, col int) (cellType CellType, err error) {
	if row < 0 || row >= grid.Rows || col < 0 || col >= grid.Cols {
		return EMPTY, fmt.Errorf("out of bound")
	}
	return grid.Data[row][col], nil
}
func (grid *Grid) canSandEnter(row int, col int) bool {
	cellType, err := grid.At(row, col)
	if err != nil || grid.moved[row][col] {
		return false
	}
	return cellType == EMPTY || cellType == WATER
}
func (grid *Grid) canWaterEnter(row int, col int) bool {
	cellType, err := grid.At(row, col)
	return err == nil && !grid.moved[row][col] && cellType == EMPTY
}
func (grid *Grid) swap(row1 int, col1 int, row2 int, col2 int) {
	if row1 < 0 || row1 >= grid.Rows || col1 < 0 || col1 >= grid.Cols {
		return
	}
	if row2 < 0 || row2 >= grid.Rows || col2 < 0 || col2 >= grid.Cols {
		return
	}

	tmp := grid.Data[row1][col1]
	grid.Data[row1][col1] = grid.Data[row2][col2]
	grid.Data[row2][col2] = tmp
	grid.moved[row1][col1] = true
	grid.moved[row2][col2] = true
}
func (grid *Grid) Draw(row int, col int, cellType CellType) {
	grid.paint(row, col, cellType)
	grid.paint(row+1, col, cellType)
	grid.paint(row-1, col, cellType)
	grid.paint(row, col+1, cellType)
	grid.paint(row, col-1, cellType)
}

func (grid *Grid) paint(row int, col int, cellType CellType) {
	if row < 0 || row >= grid.Rows || col < 0 || col >= grid.Cols {
		return
	}

	if utils.RandomBool() {
		grid.Data[row][col] = cellType
	}
}

func NewGrid(rows int, cols int, cellSize int) *Grid {
	grid := &Grid{Cols: cols, Rows: rows, CellSize: cellSize}
	grid.Init()
	return grid
}
