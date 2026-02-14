package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/thorej/go-invaders-spark/internal/render"
	"github.com/thorej/go-invaders-spark/internal/scene"
	"github.com/thorej/go-invaders-spark/internal/spritegen"
)

const (
	LogicalWidth  = 224
	LogicalHeight = 256
)

type Game struct {
	registry *render.SpriteRegistry
	gallery  *scene.Gallery
	ticks    int

	state    GameState
	progress Progress

	player      Player
	enemies     []Enemy
	projectiles []Projectile
	explosions  []Explosion
}

func New() (*Game, error) {
	sets, err := spritegen.GenerateAll(spritegen.Config{})
	if err != nil {
		return nil, fmt.Errorf("generate sprites: %w", err)
	}
	registry := render.NewSpriteRegistry(sets)
	g := &Game{
		registry:    registry,
		gallery:     scene.NewGallery(registry),
		state:       StateTitle,
		progress:    Progress{Lives: 3, Wave: 1},
		enemies:     make([]Enemy, 0, 32),
		projectiles: make([]Projectile, 0, 32),
		explosions:  make([]Explosion, 0, 32),
	}
	g.resetPlayer()
	return g, nil
}

func (g *Game) resetPlayer() {
	g.player = Player{
		Pos:      Vec2{X: LogicalWidth / 2, Y: LogicalHeight - 24},
		SpriteID: spritegen.IDPlayerShip,
		Alive:    true,
	}
}

func (g *Game) Update() error {
	g.ticks++

	switch g.state {
	case StateTitle:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.state = StatePlaying
		}
	case StatePlaying:
		g.updatePlayerInput()
		g.updateProjectilesMotion()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x04, G: 0x04, B: 0x10, A: 0xFF})
	g.gallery.Draw(screen, g.ticks)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return LogicalWidth, LogicalHeight
}
