package game

import "math/rand"

const starCount = 96

func (g *Game) initStars() {
	r := rand.New(rand.NewSource(1979))
	g.stars = make([]Star, 0, starCount)
	for i := 0; i < starCount; i++ {
		g.stars = append(g.stars, Star{
			Pos: Vec2{
				X: float64(r.Intn(LogicalWidth)),
				Y: float64(r.Intn(LogicalHeight)),
			},
			Speed: 0.15 + float64(r.Intn(40))/100.0,
			Phase: r.Intn(60),
		})
	}
}

func (g *Game) updateStars() {
	for i := range g.stars {
		g.stars[i].Pos.Y += g.stars[i].Speed
		if g.stars[i].Pos.Y >= LogicalHeight {
			g.stars[i].Pos.Y -= LogicalHeight
		}
	}
}
