package game

import "github.com/teedjayj/go-galaxians/internal/render"

type Vec2 struct {
	X float64
	Y float64
}

type EntityKind int

const (
	KindPlayer EntityKind = iota
	KindEnemy
	KindProjectile
	KindExplosion
)

type Player struct {
	Pos          Vec2
	Vel          Vec2
	SpriteID     render.SpriteID
	FrameTick    int
	Alive        bool
	FireCooldown int
}

type Enemy struct {
	Pos        Vec2
	Vel        Vec2
	SpriteID   render.SpriteID
	FrameTick  int
	Alive      bool
	Diving     bool
	Formation  Vec2
	Entering   bool
	EntryDelay int
}

type Projectile struct {
	Pos       Vec2
	Vel       Vec2
	SpriteID  render.SpriteID
	FrameTick int
	Alive     bool
	FromEnemy bool
}

type Explosion struct {
	Pos       Vec2
	SpriteID  render.SpriteID
	FrameTick int
	Alive     bool
}

type Star struct {
	Pos   Vec2
	Speed float64
	Phase int
}
