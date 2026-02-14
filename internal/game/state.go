package game

type GameState int

const (
	StateTitle GameState = iota
	StatePlaying
	StatePlayerDead
	StateWaveClear
	StateGameOver
)

type Progress struct {
	Score     int
	HighScore int
	Lives     int
	Wave      int
}
