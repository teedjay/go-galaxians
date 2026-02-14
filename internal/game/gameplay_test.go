package game

import (
	"testing"

	"github.com/thorej/go-galaxians/internal/spritegen"
)

func TestHitAABB(t *testing.T) {
	if !hitAABB(10, 10, 4, 4, 11, 11, 4, 4) {
		t.Fatalf("expected overlap")
	}
	if hitAABB(10, 10, 2, 2, 20, 20, 2, 2) {
		t.Fatalf("expected no overlap")
	}
}

func TestPlayerHitLifeAndGameOver(t *testing.T) {
	g := &Game{progress: Progress{Lives: 2}}
	g.player = Player{Alive: true}
	g.playerHit()
	if g.progress.Lives != 1 {
		t.Fatalf("expected 1 life, got %d", g.progress.Lives)
	}
	if g.state != StatePlayerDead {
		t.Fatalf("expected state player dead, got %v", g.state)
	}

	g.player.Alive = true
	g.playerHit()
	if g.progress.Lives != 0 {
		t.Fatalf("expected 0 lives, got %d", g.progress.Lives)
	}
	if g.state != StateGameOver {
		t.Fatalf("expected game over, got %v", g.state)
	}
}

func TestWaveClearState(t *testing.T) {
	g := &Game{state: StatePlaying}
	g.enemies = []Enemy{{Alive: false}, {Alive: false}}
	g.updateCombat()
	if g.state != StateWaveClear {
		t.Fatalf("expected wave clear, got %v", g.state)
	}
}

func TestUpdateDeadStateRespawns(t *testing.T) {
	g := &Game{state: StatePlayerDead, progress: Progress{Lives: 2}}
	g.stateTicks = playerRespawnTicks
	g.updateDeadState()
	if g.state != StatePlaying {
		t.Fatalf("expected playing after respawn, got %v", g.state)
	}
	if !g.player.Alive {
		t.Fatalf("expected player alive after respawn")
	}
}

func TestDeterministicDiveSelection(t *testing.T) {
	g1 := &Game{rng: newEnemyRNG(), state: StatePlaying}
	g2 := &Game{rng: newEnemyRNG(), state: StatePlaying}
	g1.setupWave()
	g2.setupWave()
	readyEnemies(g1.enemies)
	readyEnemies(g2.enemies)
	g1.diveCooldown = 0
	g2.diveCooldown = 0
	g1.player.Pos.X = 50
	g2.player.Pos.X = 50

	g1.updateDiveSelection()
	g2.updateDiveSelection()

	i1 := firstDivingIndex(g1.enemies)
	i2 := firstDivingIndex(g2.enemies)
	if i1 != i2 {
		t.Fatalf("expected same diving index, got %d and %d", i1, i2)
	}
	if i1 < 0 {
		t.Fatalf("expected one enemy to start dive")
	}
}

func TestDeterministicEnemyFire(t *testing.T) {
	g1 := &Game{rng: newEnemyRNG(), ticks: enemyFireFrames}
	g2 := &Game{rng: newEnemyRNG(), ticks: enemyFireFrames}
	g1.setupWave()
	g2.setupWave()
	readyEnemies(g1.enemies)
	readyEnemies(g2.enemies)

	g1.updateEnemyFire()
	g2.updateEnemyFire()

	if len(g1.projectiles) != 1 || len(g2.projectiles) != 1 {
		t.Fatalf("expected one projectile each")
	}
	if g1.projectiles[0].Pos.X != g2.projectiles[0].Pos.X || g1.projectiles[0].Pos.Y != g2.projectiles[0].Pos.Y {
		t.Fatalf("expected deterministic projectile spawn positions")
	}
	if g1.projectiles[0].SpriteID != g2.projectiles[0].SpriteID {
		t.Fatalf("expected deterministic projectile sprite selection")
	}
}

func TestEnemyScoreByTypeAndDive(t *testing.T) {
	if enemyScore(Enemy{SpriteID: spritegen.IDEnemyRedFlight}) != 50 {
		t.Fatalf("unexpected red score")
	}
	if enemyScore(Enemy{SpriteID: spritegen.IDEnemyFlagshipFlight}) != 150 {
		t.Fatalf("unexpected flagship score")
	}
	if enemyScore(Enemy{SpriteID: spritegen.IDEnemyFlagshipFlight, Diving: true}) != 220 {
		t.Fatalf("unexpected flagship dive score")
	}
}

func TestExtraLifeThreshold(t *testing.T) {
	g := &Game{progress: Progress{Lives: 3}, nextExtraLife: firstExtraLifeAt}
	g.addScore(firstExtraLifeAt)
	if g.progress.Lives != 4 {
		t.Fatalf("expected extra life at threshold, got %d", g.progress.Lives)
	}
	if g.nextExtraLife != firstExtraLifeAt+extraLifeStepScore {
		t.Fatalf("unexpected next threshold %d", g.nextExtraLife)
	}
}

func firstDivingIndex(enemies []Enemy) int {
	for i, e := range enemies {
		if e.Diving {
			return i
		}
	}
	return -1
}

func readyEnemies(enemies []Enemy) {
	for i := range enemies {
		enemies[i].Entering = false
		enemies[i].Pos = enemies[i].Formation
	}
}
