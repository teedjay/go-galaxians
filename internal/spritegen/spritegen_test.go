package spritegen

import (
	"testing"

	"github.com/thorej/go-invaders-spark/internal/render"
)

func TestGenerateAllExpectedIDsPresent(t *testing.T) {
	sets, err := GenerateAll(Config{})
	if err != nil {
		t.Fatalf("GenerateAll failed: %v", err)
	}

	for _, id := range BaseExpectedIDs {
		if _, ok := sets[id]; !ok {
			t.Fatalf("missing expected sprite id %q", id)
		}
	}
	for _, ch := range GlyphCharset {
		id := GlyphIDForRune(ch)
		if _, ok := sets[id]; !ok {
			t.Fatalf("missing expected glyph id %q", id)
		}
	}
}

func TestFrameCounts(t *testing.T) {
	sets, err := GenerateAll(Config{})
	if err != nil {
		t.Fatalf("GenerateAll failed: %v", err)
	}

	expected := map[render.SpriteID]int{
		IDPlayerShip:          2,
		IDPlayerExplosion:     4,
		IDEnemyRedFlight:      2,
		IDEnemyRedDive:        2,
		IDEnemyPurpleFlight:   2,
		IDEnemyPurpleDive:     2,
		IDEnemyFlagshipFlight: 2,
		IDEnemyFlagshipDive:   2,
		IDEnemyEscortFlight:   2,
		IDEnemyEscortDive:     2,
		IDBulletPlayer:        2,
		IDBulletEnemyA:        2,
		IDBulletEnemyB:        2,
		IDFxExplosionSmall:    4,
		IDFxExplosionLarge:    6,
		IDFxSparkle:           2,
	}
	for id, want := range expected {
		got := len(sets[id].Frames)
		if got != want {
			t.Fatalf("frame count mismatch for %s: got %d want %d", id, got, want)
		}
	}
}

func TestDimensionsAndOpaquePixels(t *testing.T) {
	sets, err := GenerateAll(Config{})
	if err != nil {
		t.Fatalf("GenerateAll failed: %v", err)
	}

	for id, set := range sets {
		if len(set.Frames) == 0 {
			t.Fatalf("set %s has no frames", id)
		}
		w := set.Frames[0].W
		h := set.Frames[0].H
		if w <= 0 || h <= 0 {
			t.Fatalf("set %s has invalid frame dimensions %dx%d", id, w, h)
		}
		for i, frame := range set.Frames {
			if frame.W != w || frame.H != h {
				t.Fatalf("set %s frame %d dimensions changed from %dx%d to %dx%d", id, i, w, h, frame.W, frame.H)
			}
			if frame.OpaquePixels <= 0 {
				t.Fatalf("set %s frame %d has no opaque pixels", id, i)
			}
		}
	}
}

func TestDeterministicHashes(t *testing.T) {
	a, err := GenerateAll(Config{})
	if err != nil {
		t.Fatalf("GenerateAll first run failed: %v", err)
	}
	b, err := GenerateAll(Config{})
	if err != nil {
		t.Fatalf("GenerateAll second run failed: %v", err)
	}

	if len(a) != len(b) {
		t.Fatalf("set count mismatch %d vs %d", len(a), len(b))
	}

	for id, setA := range a {
		setB, ok := b[id]
		if !ok {
			t.Fatalf("missing sprite %s in second run", id)
		}
		if len(setA.Frames) != len(setB.Frames) {
			t.Fatalf("frame count mismatch for %s", id)
		}
		for i := range setA.Frames {
			if setA.Frames[i].Hash != setB.Frames[i].Hash {
				t.Fatalf("hash mismatch for %s frame %d", id, i)
			}
		}
	}
}
