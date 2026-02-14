package game

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/teedjayj/go-galaxians/internal/render"
	"github.com/teedjayj/go-galaxians/internal/spritegen"
)

const (
	LogicalWidth       = 224
	LogicalHeight      = 256
	firstExtraLifeAt   = 5000
	extraLifeStepScore = 7000
)

type Game struct {
	registry *render.SpriteRegistry
	ticks    int

	state      GameState
	stateTicks int
	progress   Progress

	player      Player
	enemies     []Enemy
	projectiles []Projectile
	explosions  []Explosion
	stars       []Star

	rng           *rand.Rand
	diveCooldown  int
	nextExtraLife int
}

func New() (*Game, error) {
	sets, err := spritegen.GenerateAll(spritegen.Config{})
	if err != nil {
		return nil, fmt.Errorf("generate sprites: %w", err)
	}
	registry := render.NewSpriteRegistry(sets)
	g := &Game{
		registry:      registry,
		state:         StateTitle,
		progress:      Progress{Lives: 3, Wave: 1},
		enemies:       make([]Enemy, 0, 32),
		projectiles:   make([]Projectile, 0, 32),
		explosions:    make([]Explosion, 0, 32),
		rng:           newEnemyRNG(),
		nextExtraLife: firstExtraLifeAt,
	}
	g.initStars()
	g.resetPlayer()
	g.setupWave()
	return g, nil
}

func (g *Game) resetPlayer() {
	g.player = Player{
		Pos:      Vec2{X: LogicalWidth / 2, Y: LogicalHeight - 24},
		SpriteID: spritegen.IDPlayerShip,
		Alive:    true,
	}
}

func (g *Game) restartGame() {
	g.state = StatePlaying
	g.stateTicks = 0
	g.progress = Progress{Lives: 3, Wave: 1, HighScore: g.progress.HighScore}
	g.projectiles = g.projectiles[:0]
	g.explosions = g.explosions[:0]
	g.nextExtraLife = firstExtraLifeAt
	g.setupWave()
	g.resetPlayer()
}

func (g *Game) Update() error {
	g.ticks++
	g.updateStars()

	switch g.state {
	case StateTitle:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.restartGame()
		}
	case StatePlaying:
		g.stateTicks++
		g.updatePlayerInput()
		g.updateProjectilesMotion()
		g.updateEnemies()
		g.updateCombat()
	case StatePlayerDead:
		g.updateDeadState()
	case StateWaveClear:
		g.updateWaveClear()
	case StateGameOver:
		g.updateExplosions()
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.restartGame()
		}
	}
	return nil
}

func (g *Game) Layout(_, _ int) (int, int) {
	return LogicalWidth, LogicalHeight
}
