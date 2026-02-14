package game

import (
	"math"
	"math/rand"

	"github.com/teedjayj/go-galaxians/internal/render"
	"github.com/teedjayj/go-galaxians/internal/spritegen"
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
	entrySpeed       = 1.25
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

	entryGap := 5
	if g.progress.Wave > 3 {
		entryGap = 4
	}
	if g.progress.Wave > 6 {
		entryGap = 3
	}

	for row := 0; row < formationRows; row++ {
		for col := 0; col < formationCols; col++ {
			targetX := formationStartX + float64(col)*formationStepX
			targetY := formationStartY + float64(row)*formationStepY
			idx := row*formationCols + col
			startY := -40.0 - float64(row*10)
			g.enemies = append(g.enemies, Enemy{
				Pos:        Vec2{X: targetX, Y: startY},
				Vel:        Vec2{X: formationSpeed, Y: 0},
				SpriteID:   sprites[row%len(sprites)],
				Alive:      true,
				Formation:  Vec2{X: targetX, Y: targetY},
				Entering:   true,
				EntryDelay: idx * entryGap,
			})
		}
	}
}

func (g *Game) updateEnemies() {
	g.updateEnemyEntry()
	g.updateFormationMotion()
	g.updateDiveSelection()
	g.updateDiveMotion()
}

func (g *Game) updateEnemyEntry() {
	for i := range g.enemies {
		e := &g.enemies[i]
		if !e.Alive || !e.Entering {
			continue
		}
		if e.EntryDelay > 0 {
			e.EntryDelay--
			continue
		}
		e.Pos.Y += entrySpeed
		e.Pos.X += math.Sin((float64(g.ticks)+e.Formation.X)/17.0) * 0.25
		if e.Pos.Y >= e.Formation.Y {
			e.Pos = e.Formation
			e.Entering = false
			e.Vel.Y = 0
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
		if !e.Alive || e.Diving || e.Entering {
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
		if !g.enemies[i].Alive || g.enemies[i].Diving || g.enemies[i].Entering {
			continue
		}
		if flip {
			g.enemies[i].Vel.X *= -1
			g.enemies[i].Pos.Y += formationDrop
			g.enemies[i].Formation.Y += formationDrop
		}
		g.enemies[i].Pos.X += g.enemies[i].Vel.X
		g.enemies[i].Formation.X += g.enemies[i].Vel.X
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
		if g.enemies[i].Alive && !g.enemies[i].Diving && !g.enemies[i].Entering {
			candidates = append(candidates, i)
		}
	}
	if len(candidates) == 0 {
		return
	}
	idx := candidates[g.rng.Intn(len(candidates))]
	e := &g.enemies[idx]
	e.Diving = true
	e.FrameTick = 0
	vx := 0.0
	if g.player.Pos.X > e.Pos.X {
		vx = 0.9
	} else {
		vx = -0.9
	}
	e.Vel = Vec2{X: vx, Y: 1.3}
	g.diveCooldown = diveLaunchFrames
	g.audio.playDive()
}

func (g *Game) updateDiveMotion() {
	for i := range g.enemies {
		e := &g.enemies[i]
		if !e.Alive || !e.Diving {
			continue
		}
		e.FrameTick++
		if e.FrameTick < 55 {
			drift := 0.2 * math.Sin(float64(e.FrameTick)/6.0)
			e.Pos.X += e.Vel.X + drift
			e.Pos.Y += e.Vel.Y + 0.2*math.Sin(float64(e.FrameTick)/8.0)
		} else {
			dx := e.Formation.X - e.Pos.X
			dy := e.Formation.Y - e.Pos.Y
			d := math.Hypot(dx, dy)
			if d < 2.5 {
				e.Pos = e.Formation
				e.Diving = false
				e.FrameTick = 0
				e.Vel = Vec2{X: formationSpeed, Y: 0}
				continue
			}
			if d > 0 {
				e.Vel.X = (dx / d) * 1.4
				e.Vel.Y = (dy / d) * 1.4
			}
			e.Pos.X += e.Vel.X
			e.Pos.Y += e.Vel.Y
		}
	}
}
