package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"mo.example/game"
)

func main() {
	ebiten.SetWindowSize(600, 800)
	ebiten.SetWindowTitle("SandBox")
	g := &game.Game{
		Grid:            game.NewGrid(160, 120, 5),
		CurrentCellType: game.SAND,
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
