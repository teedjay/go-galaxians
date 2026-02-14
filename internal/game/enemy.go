package game

import (
	"math"
	"math/rand"

	"github.com/thorej/go-invaders-spark/internal/render"
	"github.com/thorej/go-invaders-spark/internal/spritegen"
)

const (
	formationCols    = 8
	formationRows    = 4
	formationStartX  = 28.0
	formationStartY  = 32.0
	formationStepX   = 20.0
	formationStepY   = 16.0
	formationSpeed   = 0.35
	formationDrop    = 8.0
	diveLaunchFrames = 90
)

func newEnemyRNG() *rand.Rand {
	return rand.New(rand.NewSource(424242))
}

func (g *Game) setupWave() {
	g.enemies = g.enemies[:0]
	g.diveCooldown = diveLaunchFrames
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

func (g *Game) updateEnemies() {
	g.updateFormationMotion()
	g.updateDiveSelection()
	g.updateDiveMotion()
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

func (g *Game) updateDiveSelection() {
	if g.state != StatePlaying {
		return
	}
	if g.diveCooldown > 0 {
		g.diveCooldown--
		return
	}
	candidates := make([]int, 0, len(g.enemies))
	for i := range g.enemies {
		if g.enemies[i].Alive && !g.enemies[i].Diving {
			candidates = append(candidates, i)
		}
	}
	if len(candidates) == 0 {
		return
	}
	idx := candidates[g.rng.Intn(len(candidates))]
	e := &g.enemies[idx]
	e.Diving = true
	vx := 0.0
	if g.player.Pos.X > e.Pos.X {
		vx = 0.7
	} else {
		vx = -0.7
	}
	e.Vel = Vec2{X: vx, Y: 1.2}
	g.diveCooldown = diveLaunchFrames
}

func (g *Game) updateDiveMotion() {
	for i := range g.enemies {
		e := &g.enemies[i]
		if !e.Alive || !e.Diving {
			continue
		}
		e.Pos.X += e.Vel.X
		e.Pos.Y += e.Vel.Y

		if e.Pos.Y > LogicalHeight*0.65 {
			dx := e.Formation.X - e.Pos.X
			dy := e.Formation.Y - e.Pos.Y
			d := math.Hypot(dx, dy)
			if d < 2.5 {
				e.Pos = e.Formation
				e.Diving = false
				e.Vel = Vec2{X: formationSpeed, Y: 0}
				continue
			}
			if d > 0 {
				e.Vel.X = (dx / d) * 1.2
				e.Vel.Y = (dy / d) * 1.2
			}
		}
	}
}
