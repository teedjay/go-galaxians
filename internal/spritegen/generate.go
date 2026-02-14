package spritegen

import (
	"fmt"
	"image"

	"github.com/teedjayj/go-galaxians/internal/render"
)

type Config struct{}

func GenerateAll(cfg Config) (map[render.SpriteID]render.SpriteSet, error) {
	_ = cfg
	sets := make(map[render.SpriteID]render.SpriteSet)

	add := func(id render.SpriteID, duration int, masks [][]string) error {
		frames := make([]render.Frame, 0, len(masks))
		for _, mask := range masks {
			frame, err := frameFromMask(mask, arcadePalette, image.Point{X: len(mask[0]) / 2, Y: len(mask) / 2})
			if err != nil {
				return fmt.Errorf("build %s: %w", id, err)
			}
			frames = append(frames, frame)
		}
		sets[id] = render.SpriteSet{ID: id, Frames: frames, FrameDuration: duration}
		return nil
	}

	for _, spec := range allSpecs() {
		if err := add(spec.id, spec.frameDuration, spec.frames); err != nil {
			return nil, err
		}
	}

	glyphSpecs := glyphSpecs()
	for _, spec := range glyphSpecs {
		if err := add(spec.id, spec.frameDuration, spec.frames); err != nil {
			return nil, err
		}
	}

	return sets, nil
}
