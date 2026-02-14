package scene

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/teedjayj/go-galaxians/internal/render"
	"github.com/teedjayj/go-galaxians/internal/spritegen"
)

func DrawText(reg *render.SpriteRegistry, dst *ebiten.Image, text string, x, y int, scale float64) {
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
