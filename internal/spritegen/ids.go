package spritegen

import "github.com/thorej/go-galaxians/internal/render"

const (
	IDPlayerShip      render.SpriteID = "player.ship"
	IDPlayerExplosion render.SpriteID = "player.explosion"

	IDEnemyRedFlight      render.SpriteID = "enemy.red.flight"
	IDEnemyRedDive        render.SpriteID = "enemy.red.dive"
	IDEnemyPurpleFlight   render.SpriteID = "enemy.purple.flight"
	IDEnemyPurpleDive     render.SpriteID = "enemy.purple.dive"
	IDEnemyFlagshipFlight render.SpriteID = "enemy.flagship.flight"
	IDEnemyFlagshipDive   render.SpriteID = "enemy.flagship.dive"
	IDEnemyEscortFlight   render.SpriteID = "enemy.escort.flight"
	IDEnemyEscortDive     render.SpriteID = "enemy.escort.dive"

	IDBulletPlayer render.SpriteID = "bullet.player"
	IDBulletEnemyA render.SpriteID = "bullet.enemy_a"
	IDBulletEnemyB render.SpriteID = "bullet.enemy_b"

	IDFxExplosionSmall render.SpriteID = "fx.explosion_small"
	IDFxExplosionLarge render.SpriteID = "fx.explosion_large"
	IDFxSparkle        render.SpriteID = "fx.sparkle"
)

var GlyphCharset = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ-:")

var BaseExpectedIDs = []render.SpriteID{
	IDPlayerShip,
	IDPlayerExplosion,
	IDEnemyRedFlight,
	IDEnemyRedDive,
	IDEnemyPurpleFlight,
	IDEnemyPurpleDive,
	IDEnemyFlagshipFlight,
	IDEnemyFlagshipDive,
	IDEnemyEscortFlight,
	IDEnemyEscortDive,
	IDBulletPlayer,
	IDBulletEnemyA,
	IDBulletEnemyB,
	IDFxExplosionSmall,
	IDFxExplosionLarge,
	IDFxSparkle,
}
