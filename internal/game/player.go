package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/teedjayj/go-galaxians/internal/spritegen"
)

const (
	playerSpeed      = 2.0
	playerShotSpeed  = -3.5
	playerShotMax    = 2
	playerFireFrames = 10
)

func (g *Game) updatePlayerInput() {
	if !g.player.Alive {
		return
	}

	vx := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		vx -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		vx += playerSpeed
	}
	g.player.Vel.X = vx
	g.player.Pos.X += g.player.Vel.X

	if g.player.Pos.X < 8 {
		g.player.Pos.X = 8
	}
	if g.player.Pos.X > LogicalWidth-8 {
		g.player.Pos.X = LogicalWidth - 8
	}

	if g.player.FireCooldown > 0 {
		g.player.FireCooldown--
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.player.FireCooldown == 0 && g.countPlayerShots() < playerShotMax {
		g.spawnPlayerShot()
		g.player.FireCooldown = playerFireFrames
	}
}

func (g *Game) countPlayerShots() int {
	count := 0
	for _, p := range g.projectiles {
		if p.Alive && !p.FromEnemy {
			count++
		}
	}
	return count
}

func (g *Game) spawnPlayerShot() {
	g.projectiles = append(g.projectiles, Projectile{
		Pos:       Vec2{X: g.player.Pos.X, Y: g.player.Pos.Y - 8},
		Vel:       Vec2{X: 0, Y: playerShotSpeed},
		SpriteID:  spritegen.IDBulletPlayer,
		Alive:     true,
		FromEnemy: false,
	})
}

func (g *Game) updateProjectilesMotion() {
	for i := range g.projectiles {
		if !g.projectiles[i].Alive {
			continue
		}
		g.projectiles[i].Pos.X += g.projectiles[i].Vel.X
		g.projectiles[i].Pos.Y += g.projectiles[i].Vel.Y
		if g.projectiles[i].Pos.Y < -10 || g.projectiles[i].Pos.Y > LogicalHeight+10 {
			g.projectiles[i].Alive = false
		}
	}
}
