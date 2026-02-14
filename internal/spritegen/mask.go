package spritegen

import (
	"errors"
	"fmt"
	"hash/fnv"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/thorej/go-invaders-spark/internal/render"
)

func frameFromMask(mask []string, palette map[rune]color.RGBA, anchor image.Point) (render.Frame, error) {
	if len(mask) == 0 {
		return render.Frame{}, errors.New("mask cannot be empty")
	}
	w := len(mask[0])
	if w == 0 {
		return render.Frame{}, errors.New("mask width cannot be zero")
	}
	for i := 1; i < len(mask); i++ {
		if len(mask[i]) != w {
			return render.Frame{}, fmt.Errorf("inconsistent row width at row %d", i)
		}
	}

	h := len(mask)
	pix := make([]byte, w*h*4)
	opaque := 0
	hasher := fnv.New64a()

	for y, row := range mask {
		for x, ch := range row {
			col, ok := palette[ch]
			if !ok {
				return render.Frame{}, fmt.Errorf("unknown palette rune %q", ch)
			}
			if col.A > 0 {
				opaque++
			}
			o := (y*w + x) * 4
			pix[o+0] = col.R
			pix[o+1] = col.G
			pix[o+2] = col.B
			pix[o+3] = col.A
		}
	}

	if _, err := hasher.Write(pix); err != nil {
		return render.Frame{}, err
	}

	img := ebiten.NewImage(w, h)
	img.WritePixels(pix)

	return render.Frame{
		Image:        img,
		W:            w,
		H:            h,
		Anchor:       anchor,
		OpaquePixels: opaque,
		Hash:         hasher.Sum64(),
	}, nil
}
