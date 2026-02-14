package scene

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/thorej/go-invaders-spark/internal/render"
	"github.com/thorej/go-invaders-spark/internal/spritegen"
)

type Gallery struct {
	registry *render.SpriteRegistry
	ids      []render.SpriteID
}

func NewGallery(registry *render.SpriteRegistry) *Gallery {
	all := registry.List("")
	ids := make([]render.SpriteID, 0, len(all))
	for _, id := range all {
		if strings.HasPrefix(string(id), "ui.glyph.") {
			continue
		}
		ids = append(ids, id)
	}
	return &Gallery{registry: registry, ids: ids}
}

func (g *Gallery) Draw(screen *ebiten.Image, ticks int) {
	drawText(g.registry, screen, fmt.Sprintf("SETS:%d FRAMES:%d", g.registry.Count(), g.registry.TotalFrames()), 4, 4, 1)
	drawText(g.registry, screen, "GALAXIANS SPRITE GALLERY", 4, 12, 1)

	const (
		cols      = 4
		cellW     = 56
		cellH     = 54
		startX    = 2
		startY    = 24
		labelY    = 30
		spriteTop = 2
	)

	for i, id := range g.ids {
		set := g.registry.MustGet(id)
		if len(set.Frames) == 0 {
			continue
		}
		col := i % cols
		row := i / cols
		x := startX + col*cellW
		y := startY + row*cellH

		frameIndex := 0
		if len(set.Frames) > 1 && set.FrameDuration > 0 {
			frameIndex = (ticks / set.FrameDuration) % len(set.Frames)
		}
		frame := set.Frames[frameIndex]

		scale := 2.0
		if frame.W > 10 || frame.H > 10 {
			scale = 1.5
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(float64(x+2), float64(y+spriteTop))
		screen.DrawImage(frame.Image, op)

		drawText(g.registry, screen, strings.ToUpper(string(id)), x+1, y+labelY, 0.5)
	}

	drawGlyphPreview(g.registry, screen)
}

func drawGlyphPreview(reg *render.SpriteRegistry, screen *ebiten.Image) {
	drawText(reg, screen, "UI GLYPHS", 4, 246, 1)
	x := 64
	y := 246
	for _, ch := range spritegen.GlyphCharset {
		drawText(reg, screen, string(ch), x, y, 1)
		x += 6
		if x > 216 {
			break
		}
	}
}

func drawText(reg *render.SpriteRegistry, dst *ebiten.Image, text string, x, y int, scale float64) {
	cursorX := x
	for _, raw := range text {
		if raw == ' ' {
			cursorX += int(4 * scale)
			continue
		}
		r := raw
		if raw >= 'a' && raw <= 'z' {
			r = raw - ('a' - 'A')
		}
		id := spritegen.GlyphIDForRune(r)
		set, ok := reg.Get(id)
		if !ok || len(set.Frames) == 0 {
			cursorX += int(4 * scale)
			continue
		}
		frame := set.Frames[0]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(float64(cursorX), float64(y))
		dst.DrawImage(frame.Image, op)
		cursorX += int(float64(frame.W+1) * scale)
	}
}
