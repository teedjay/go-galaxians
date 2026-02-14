package render

import (
	"image"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteID string

type Frame struct {
	Image        *ebiten.Image
	W            int
	H            int
	Anchor       image.Point
	OpaquePixels int
	Hash         uint64
}

type SpriteSet struct {
	ID            SpriteID
	Frames        []Frame
	FrameDuration int
}

type Registry interface {
	Get(id SpriteID) (SpriteSet, bool)
	MustGet(id SpriteID) SpriteSet
	List(prefix string) []SpriteID
	Count() int
	TotalFrames() int
}

type SpriteRegistry struct {
	sets map[SpriteID]SpriteSet
	ids  []SpriteID
}

func NewSpriteRegistry(sets map[SpriteID]SpriteSet) *SpriteRegistry {
	ids := make([]SpriteID, 0, len(sets))
	for id := range sets {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	return &SpriteRegistry{sets: sets, ids: ids}
}

func (r *SpriteRegistry) Get(id SpriteID) (SpriteSet, bool) {
	set, ok := r.sets[id]
	return set, ok
}

func (r *SpriteRegistry) MustGet(id SpriteID) SpriteSet {
	set, ok := r.Get(id)
	if !ok {
		panic("missing sprite set: " + string(id))
	}
	return set
}

func (r *SpriteRegistry) List(prefix string) []SpriteID {
	if prefix == "" {
		out := make([]SpriteID, len(r.ids))
		copy(out, r.ids)
		return out
	}
	out := make([]SpriteID, 0)
	for _, id := range r.ids {
		if len(id) >= len(prefix) && string(id[:len(prefix)]) == prefix {
			out = append(out, id)
		}
	}
	return out
}

func (r *SpriteRegistry) Count() int {
	return len(r.sets)
}

func (r *SpriteRegistry) TotalFrames() int {
	total := 0
	for _, set := range r.sets {
		total += len(set.Frames)
	}
	return total
}
