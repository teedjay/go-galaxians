package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/thorej/go-invaders-spark/internal/game"
)

func main() {
	ebiten.SetWindowSize(game.LogicalWidth*3, game.LogicalHeight*3)
	ebiten.SetWindowTitle("Go Galaxians - Sprite Gallery")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g, err := game.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
