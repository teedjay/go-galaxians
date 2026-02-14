package game

import (
	"github.com/thorej/go-galaxians/internal/render"
	"github.com/thorej/go-galaxians/internal/spritegen"
)

const (
	enemyFireFrames    = 45
	playerRespawnTicks = 60
)

func (g *Game) updateCombat() {
	g.updateEnemyFire()
	g.checkCollisions()
	g.cleanupDeadProjectiles()
	g.updateExplosions()
	if g.aliveEnemyCount() == 0 {
		g.state = StateWaveClear
		g.stateTicks = 0
	}
}

func (g *Game) updateEnemyFire() {
	if g.ticks%enemyFireFrames != 0 {
		return
	}
	candidates := make([]int, 0, len(g.enemies))
	for i := range g.enemies {
		if g.enemies[i].Alive && !g.enemies[i].Entering {
			candidates = append(candidates, i)
		}
	}
	if len(candidates) == 0 {
		return
	}
	e := g.enemies[candidates[g.rng.Intn(len(candidates))]]
	id := spritegen.IDBulletEnemyA
	if g.rng.Intn(2) == 1 {
		id = spritegen.IDBulletEnemyB
	}
	g.projectiles = append(g.projectiles, Projectile{
		Pos:       Vec2{X: e.Pos.X, Y: e.Pos.Y + 6},
		Vel:       Vec2{X: 0, Y: 2.0},
		SpriteID:  id,
		Alive:     true,
		FromEnemy: true,
	})
}

func (g *Game) checkCollisions() {
	for pi := range g.projectiles {
		p := &g.projectiles[pi]
		if !p.Alive {
			continue
		}
		if p.FromEnemy {
			if g.player.Alive && hitAABB(p.Pos.X, p.Pos.Y, 2, 6, g.player.Pos.X, g.player.Pos.Y, 12, 8) {
				p.Alive = false
				g.playerHit()
			}
			continue
		}

		for ei := range g.enemies {
			e := &g.enemies[ei]
			if !e.Alive {
				continue
			}
			if hitAABB(p.Pos.X, p.Pos.Y, 2, 6, e.Pos.X, e.Pos.Y, 10, 8) {
				p.Alive = false
				e.Alive = false
				g.addScore(enemyScore(*e))
				g.spawnExplosion(e.Pos, spritegen.IDFxExplosionSmall)
				break
			}
		}
	}
}

func (g *Game) addScore(points int) {
	g.progress.Score += points
	if g.progress.Score > g.progress.HighScore {
		g.progress.HighScore = g.progress.Score
	}
	for g.progress.Score >= g.nextExtraLife {
		g.progress.Lives++
		g.nextExtraLife += extraLifeStepScore
	}
}

func enemyScore(e Enemy) int {
	switch e.SpriteID {
	case spritegen.IDEnemyRedFlight:
		if e.Diving {
			return 100
		}
		return 50
	case spritegen.IDEnemyPurpleFlight:
		if e.Diving {
			return 130
		}
		return 80
	case spritegen.IDEnemyEscortFlight:
		if e.Diving {
			return 170
		}
		return 100
	case spritegen.IDEnemyFlagshipFlight:
		if e.Diving {
			return 220
		}
		return 150
	default:
		return 100
	}
}

func (g *Game) playerHit() {
	g.player.Alive = false
	g.state = StatePlayerDead
	g.stateTicks = 0
	g.spawnExplosion(g.player.Pos, spritegen.IDPlayerExplosion)
	g.progress.Lives--
	if g.progress.Lives <= 0 {
		g.state = StateGameOver
	}
}

func (g *Game) updateDeadState() {
	g.stateTicks++
	g.updateExplosions()
	if g.state != StatePlayerDead {
		return
	}
	if g.stateTicks < playerRespawnTicks {
		return
	}
	if g.progress.Lives <= 0 {
		g.state = StateGameOver
		return
	}
	g.resetPlayer()
	g.state = StatePlaying
	g.stateTicks = 0
}

func (g *Game) updateWaveClear() {
	g.stateTicks++
	if g.stateTicks < 90 {
		return
	}
	g.progress.Wave++
	g.setupWave()
	g.resetPlayer()
	g.state = StatePlaying
	g.stateTicks = 0
}

func (g *Game) spawnExplosion(pos Vec2, spriteID render.SpriteID) {
	g.explosions = append(g.explosions, Explosion{
		Pos:      pos,
		SpriteID: spriteID,
		Alive:    true,
	})
}

func (g *Game) updateExplosions() {
	for i := range g.explosions {
		if !g.explosions[i].Alive {
			continue
		}
		g.explosions[i].FrameTick++
		if g.explosions[i].FrameTick > 30 {
			g.explosions[i].Alive = false
		}
	}
}

func (g *Game) cleanupDeadProjectiles() {
	live := g.projectiles[:0]
	for _, p := range g.projectiles {
		if p.Alive {
			live = append(live, p)
		}
	}
	g.projectiles = live
}

func (g *Game) aliveEnemyCount() int {
	count := 0
	for _, e := range g.enemies {
		if e.Alive {
			count++
		}
	}
	return count
}

func hitAABB(ax, ay float64, aw, ah float64, bx, by float64, bw, bh float64) bool {
	ax0, ay0 := ax-aw/2, ay-ah/2
	ax1, ay1 := ax+aw/2, ay+ah/2
	bx0, by0 := bx-bw/2, by-bh/2
	bx1, by1 := bx+bw/2, by+bh/2
	return ax0 <= bx1 && ax1 >= bx0 && ay0 <= by1 && ay1 >= by0
}
