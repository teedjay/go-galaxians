package game

import (
	"github.com/thorej/go-invaders-spark/internal/render"
	"github.com/thorej/go-invaders-spark/internal/spritegen"
)

const (
	formationCols   = 8
	formationRows   = 4
	formationStartX = 28.0
	formationStartY = 32.0
	formationStepX  = 20.0
	formationStepY  = 16.0
	formationSpeed  = 0.35
	formationDrop   = 8.0
)

func (g *Game) setupWave() {
	g.enemies = g.enemies[:0]
	sprites := []render.SpriteID{
		spritegen.IDEnemyRedFlight,
		spritegen.IDEnemyPurpleFlight,
		spritegen.IDEnemyEscortFlight,
		spritegen.IDEnemyFlagshipFlight,
	}

	for row := 0; row < formationRows; row++ {
		for col := 0; col < formationCols; col++ {
			x := formationStartX + float64(col)*formationStepX
			y := formationStartY + float64(row)*formationStepY
			g.enemies = append(g.enemies, Enemy{
				Pos:       Vec2{X: x, Y: y},
				Vel:       Vec2{X: formationSpeed, Y: 0},
				SpriteID:  sprites[row%len(sprites)],
				Alive:     true,
				Formation: Vec2{X: x, Y: y},
			})
		}
	}
}

func (g *Game) updateFormationMotion() {
	if len(g.enemies) == 0 {
		return
	}
	left := 9999.0
	right := -9999.0
	for _, e := range g.enemies {
		if !e.Alive || e.Diving {
			continue
		}
		if e.Pos.X < left {
			left = e.Pos.X
		}
		if e.Pos.X > right {
			right = e.Pos.X
		}
	}

	flip := false
	if left < 10 || right > LogicalWidth-10 {
		flip = true
	}
	for i := range g.enemies {
		if !g.enemies[i].Alive || g.enemies[i].Diving {
			continue
		}
		if flip {
			g.enemies[i].Vel.X *= -1
			g.enemies[i].Pos.Y += formationDrop
		}
		g.enemies[i].Pos.X += g.enemies[i].Vel.X
	}
}
