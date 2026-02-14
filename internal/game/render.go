package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/thorej/go-galaxians/internal/render"
	"github.com/thorej/go-galaxians/internal/scene"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x04, G: 0x04, B: 0x10, A: 0xFF})
	g.drawStars(screen)
	g.drawHUD(screen)
	g.drawWorld(screen)

	switch g.state {
	case StateTitle:
		scene.DrawText(g.registry, screen, "GALAXIANS", 70, 96, 2)
		scene.DrawText(g.registry, screen, "PRESS SPACE", 64, 124, 1)
	case StateWaveClear:
		scene.DrawText(g.registry, screen, "WAVE CLEAR", 66, 120, 1)
	case StateGameOver:
		scene.DrawText(g.registry, screen, "GAME OVER", 68, 120, 1)
		scene.DrawText(g.registry, screen, "PRESS SPACE", 64, 138, 1)
	}
}

func (g *Game) drawStars(screen *ebiten.Image) {
	for i := range g.stars {
		star := g.stars[i]
		x := int(star.Pos.X)
		y := int(star.Pos.Y)
		if x < 0 || x >= LogicalWidth || y < 0 || y >= LogicalHeight {
			continue
		}
		twinkle := (g.ticks+star.Phase)%48 < 10
		if twinkle {
			screen.Set(x, y, color.RGBA{R: 0x7F, G: 0xE8, B: 0xFF, A: 0xFF})
		} else {
			screen.Set(x, y, color.RGBA{R: 0x8C, G: 0x8C, B: 0xAA, A: 0xFF})
		}
	}
}

func (g *Game) drawHUD(screen *ebiten.Image) {
	scene.DrawText(g.registry, screen, fmt.Sprintf("SCORE %05d", g.progress.Score), 4, 4, 1)
	scene.DrawText(g.registry, screen, fmt.Sprintf("HI-SCORE %05d", g.progress.HighScore), 78, 4, 1)
	scene.DrawText(g.registry, screen, fmt.Sprintf("LIVES %d", g.progress.Lives), 4, 14, 1)
	scene.DrawText(g.registry, screen, fmt.Sprintf("WAVE %d", g.progress.Wave), 86, 14, 1)
}

func (g *Game) drawWorld(screen *ebiten.Image) {
	if g.player.Alive {
		g.drawAnimatedSprite(screen, g.player.SpriteID, g.player.Pos, 0, 16)
	}
	for _, e := range g.enemies {
		if !e.Alive {
			continue
		}
		g.drawAnimatedSprite(screen, e.SpriteID, e.Pos, g.ticks, 18)
	}
	for _, p := range g.projectiles {
		if !p.Alive {
			continue
		}
		g.drawAnimatedSprite(screen, p.SpriteID, p.Pos, g.ticks, 8)
	}
	for _, fx := range g.explosions {
		if !fx.Alive {
			continue
		}
		g.drawAnimatedSprite(screen, fx.SpriteID, fx.Pos, fx.FrameTick, 6)
	}
}

func (g *Game) drawAnimatedSprite(screen *ebiten.Image, id render.SpriteID, pos Vec2, tick int, fallbackDur int) {
	set, ok := g.registry.Get(id)
	if !ok || len(set.Frames) == 0 {
		return
	}
	dur := set.FrameDuration
	if dur <= 0 {
		dur = fallbackDur
	}
	frameIndex := 0
	if len(set.Frames) > 1 {
		frameIndex = (tick / dur) % len(set.Frames)
	}
	frame := set.Frames[frameIndex]
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pos.X-float64(frame.W)/2, pos.Y-float64(frame.H)/2)
	screen.DrawImage(frame.Image, op)
}
