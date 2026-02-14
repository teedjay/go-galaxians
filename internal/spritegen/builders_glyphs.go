package spritegen

import (
	"fmt"
	"strings"

	"github.com/thorej/go-galaxians/internal/render"
)

var glyphMasks = map[rune][]string{
	'0': {".GG.", "G..G", "G..G", "G..G", ".GG."},
	'1': {"..G.", ".GG.", "..G.", "..G.", ".GGG"},
	'2': {".GG.", "G..G", "..G.", ".G..", "GGGG"},
	'3': {"GGG.", "...G", ".GG.", "...G", "GGG."},
	'4': {"G..G", "G..G", "GGGG", "...G", "...G"},
	'5': {"GGGG", "G...", "GGG.", "...G", "GGG."},
	'6': {".GG.", "G...", "GGG.", "G..G", ".GG."},
	'7': {"GGGG", "...G", "..G.", ".G..", "G..."},
	'8': {".GG.", "G..G", ".GG.", "G..G", ".GG."},
	'9': {".GG.", "G..G", ".GGG", "...G", ".GG."},
	'A': {".GG.", "G..G", "GGGG", "G..G", "G..G"},
	'B': {"GGG.", "G..G", "GGG.", "G..G", "GGG."},
	'C': {".GGG", "G...", "G...", "G...", ".GGG"},
	'D': {"GGG.", "G..G", "G..G", "G..G", "GGG."},
	'E': {"GGGG", "G...", "GGG.", "G...", "GGGG"},
	'F': {"GGGG", "G...", "GGG.", "G...", "G..."},
	'G': {".GGG", "G...", "G.GG", "G..G", ".GG."},
	'H': {"G..G", "G..G", "GGGG", "G..G", "G..G"},
	'I': {"GGG", ".G.", ".G.", ".G.", "GGG"},
	'J': {"..GG", "...G", "...G", "G..G", ".GG."},
	'K': {"G..G", "G.G.", "GG..", "G.G.", "G..G"},
	'L': {"G...", "G...", "G...", "G...", "GGGG"},
	'M': {"G...G", "GG.GG", "G.G.G", "G...G", "G...G"},
	'N': {"G..G", "GG.G", "G.GG", "G..G", "G..G"},
	'O': {".GG.", "G..G", "G..G", "G..G", ".GG."},
	'P': {"GGG.", "G..G", "GGG.", "G...", "G..."},
	'Q': {".GG.", "G..G", "G..G", "G.GG", ".GGG"},
	'R': {"GGG.", "G..G", "GGG.", "G.G.", "G..G"},
	'S': {".GGG", "G...", ".GG.", "...G", "GGG."},
	'T': {"GGGG", ".GG.", ".GG.", ".GG.", ".GG."},
	'U': {"G..G", "G..G", "G..G", "G..G", ".GG."},
	'V': {"G...G", "G...G", "G...G", ".G.G.", "..G.."},
	'W': {"G...G", "G...G", "G.G.G", "GG.GG", "G...G"},
	'X': {"G...G", ".G.G.", "..G..", ".G.G.", "G...G"},
	'Y': {"G...G", ".G.G.", "..G..", "..G..", "..G.."},
	'Z': {"GGGG", "...G", "..G.", ".G..", "GGGG"},
	'-': {"....", "....", "GGGG", "....", "...."},
	':': {".", "G", ".", "G", "."},
}

func glyphSpecs() []spriteSpec {
	out := make([]spriteSpec, 0, len(GlyphCharset))
	for _, ch := range GlyphCharset {
		mask, ok := glyphMasks[ch]
		if !ok {
			continue
		}
		out = append(out, spriteSpec{
			id:            GlyphIDForRune(ch),
			frameDuration: 1,
			frames:        [][]string{mask},
		})
	}
	return out
}

func GlyphIDForRune(r rune) render.SpriteID {
	r = rune(strings.ToUpper(string(r))[0])
	switch r {
	case '-':
		return "ui.glyph.dash"
	case ':':
		return "ui.glyph.colon"
	default:
		return render.SpriteID(fmt.Sprintf("ui.glyph.%c", r))
	}
}
