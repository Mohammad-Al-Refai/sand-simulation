package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type CellColor int

var (
	SAND_COLOR1  = color.RGBA{232, 196, 122, 255}
	SAND_COLOR2  = color.RGBA{214, 172, 98, 255}
	WATER_COLOR1 = color.RGBA{52, 128, 220, 255}
	WATER_COLOR2 = color.RGBA{72, 150, 235, 255}
	WALL_COLOR   = color.RGBA{0, 100, 100, 255}
)

type Game struct {
	CurrentCellType CellType
	Grid            *Grid
}

func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Cells: %d", g.Grid.Rows*g.Grid.Cols), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Current: %s", g.CurrentCellType), 0, 20)

	cs := g.Grid.CellSize
	for row := 0; row < g.Grid.Rows; row++ {
		for col := 0; col < g.Grid.Cols; col++ {
			current := g.Grid.Data[row][col]
			if current == EMPTY {
				continue
			}
			cellColor := color.RGBA{}
			switch current {
			case SAND:
				cellColor = shade(row, col, SAND_COLOR1, SAND_COLOR2)
			case WATER:
				cellColor = shade(row, col, WATER_COLOR1, WATER_COLOR2)
			case WALL:
				cellColor = WALL_COLOR
			}
			vector.FillRect(screen,
				float32(col*cs), float32(row*cs),
				float32(cs), float32(cs),
				cellColor, false)
		}
	}

}

// shade picks a color from the cell's position, so it doesn't flicker every frame
func shade(row int, col int, a color.RGBA, b color.RGBA) color.RGBA {
	if (row*73856093^col*19349663)%3 == 0 {
		return b
	}
	return a
}

func onMouseClick(mouse ebiten.MouseButton, callback func(mx int, my int)) {
	if ebiten.IsMouseButtonPressed(mouse) {
		x, y := ebiten.CursorPosition()
		callback(x, y)
	}
}

// Layout implements [ebiten.Game].
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 600, 800
}

// Update implements [ebiten.Game].
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.CurrentCellType = SAND
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		g.CurrentCellType = WATER
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		g.CurrentCellType = WALL
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Grid.Reset()
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("Exit")
	}

	onMouseClick(ebiten.MouseButtonLeft, func(x int, y int) {
		row, col, err := g.Grid.IndexAt(x, y)
		if err != nil {
			return
		}
		g.Grid.Draw(row, col, g.CurrentCellType)

	})
	onMouseClick(ebiten.MouseButtonRight, func(x int, y int) {
		row, col, err := g.Grid.IndexAt(x, y)
		if err != nil {
			return
		}
		g.Grid.Draw(row, col, EMPTY)
	})
	g.Grid.Step()
	return nil
}
