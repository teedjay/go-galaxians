package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/teedjayj/go-galaxians/internal/game"
)

func main() {
	ebiten.SetWindowSize(game.LogicalWidth*3, game.LogicalHeight*3)
	ebiten.SetWindowTitle("Go Galaxians")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g, err := game.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
