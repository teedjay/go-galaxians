package spritegen

import "github.com/teedjayj/go-galaxians/internal/render"

type spriteSpec struct {
	id            render.SpriteID
	frameDuration int
	frames        [][]string
}

func allSpecs() []spriteSpec {
	out := make([]spriteSpec, 0)
	out = append(out, playerSpecs()...)
	out = append(out, enemySpecs()...)
	out = append(out, bulletSpecs()...)
	out = append(out, effectSpecs()...)
	return out
}

func playerSpecs() []spriteSpec {
	shipFrameA := []string{
		"....C.C....",
		"...CCCCC...",
		"..CCCCC.C..",
		".CCCCCCCCC.",
		"CCCCCC.CCCC",
		".CCYYYYYCC.",
		"..Y..Y..Y..",
		"...Y...Y...",
	}
	shipFrameB := []string{
		"....C.C....",
		"...CCCCC...",
		"..CCCCC.C..",
		".CCCCCCCCC.",
		"CCCCCC.CCCC",
		".CCYYYYYCC.",
		"..YO.Y.OY..",
		"...Y...Y...",
	}
	playerExplosion := [][]string{
		{
			".....O.....",
			"..O..O..O..",
			"...ORRO....",
			".OORRYYOO..",
			"..ORYYYRO..",
			"...ORRO....",
			"..O..O..O..",
			".....O.....",
		},
		{
			"...O...O...",
			".O..R.R..O.",
			"..ORRYRRO..",
			".ORYYYYYRO.",
			"..ORRYRRO..",
			".O..R.R..O.",
			"...O...O...",
			".....O.....",
		},
		{
			"..O.....O..",
			".O..R.R..O.",
			"..R.....R..",
			"...R.Y.R...",
			"..R.....R..",
			".O..R.R..O.",
			"..O.....O..",
			".....O.....",
		},
		{
			"...........",
			"....R.R....",
			".....R.....",
			"....R.R....",
			".....R.....",
			"....R.R....",
			"...........",
			"...........",
		},
	}

	return []spriteSpec{
		{id: IDPlayerShip, frameDuration: 16, frames: [][]string{shipFrameA, shipFrameB}},
		{id: IDPlayerExplosion, frameDuration: 8, frames: playerExplosion},
	}
}

func enemySpecs() []spriteSpec {
	redFlight := [][]string{
		{
			"..R...R..",
			"...RRR...",
			".RRRRRRR.",
			"RRRRYRRRR",
			".RRRRRRR.",
			"..R...R..",
			".R.....R.",
			"R.......R",
		},
		{
			".R.....R.",
			"..RR.RR..",
			"RRRRRRRRR",
			"RRRRYRRRR",
			".RRRRRRR.",
			"..R...R..",
			"...R.R...",
			"....R....",
		},
	}
	redDive := [][]string{
		{
			"....R....",
			"...RRR...",
			"..RRYRR..",
			".RRRRRRR.",
			"RR.RRR.RR",
			"...R.R...",
			"..R...R..",
			".R.....R.",
		},
		{
			"....R....",
			"...RRR...",
			"..RYYYR..",
			".RRRRRRR.",
			"RRR.R.RRR",
			"..R...R..",
			"...R.R...",
			"....R....",
		},
	}

	purpleFlight := recolorSpecs(redFlight, 'R', 'P')
	purpleDive := recolorSpecs(redDive, 'R', 'P')
	flagshipFlight := recolorSpecs(redFlight, 'R', 'Y')
	flagshipDive := recolorSpecs(redDive, 'R', 'Y')
	escortFlight := recolorSpecs(redFlight, 'R', 'G')
	escortDive := recolorSpecs(redDive, 'R', 'G')

	return []spriteSpec{
		{id: IDEnemyRedFlight, frameDuration: 18, frames: redFlight},
		{id: IDEnemyRedDive, frameDuration: 18, frames: redDive},
		{id: IDEnemyPurpleFlight, frameDuration: 18, frames: purpleFlight},
		{id: IDEnemyPurpleDive, frameDuration: 18, frames: purpleDive},
		{id: IDEnemyFlagshipFlight, frameDuration: 18, frames: flagshipFlight},
		{id: IDEnemyFlagshipDive, frameDuration: 18, frames: flagshipDive},
		{id: IDEnemyEscortFlight, frameDuration: 18, frames: escortFlight},
		{id: IDEnemyEscortDive, frameDuration: 18, frames: escortDive},
	}
}

func bulletSpecs() []spriteSpec {
	return []spriteSpec{
		{
			id:            IDBulletPlayer,
			frameDuration: 8,
			frames: [][]string{
				{"W", "W", "W", "W"},
				{"W", "Y", "W", "Y"},
			},
		},
		{
			id:            IDBulletEnemyA,
			frameDuration: 8,
			frames: [][]string{
				{"R", "O", "R", "O"},
				{"O", "R", "O", "R"},
			},
		},
		{
			id:            IDBulletEnemyB,
			frameDuration: 8,
			frames: [][]string{
				{"P", "G", "P", "G"},
				{"G", "P", "G", "P"},
			},
		},
	}
}

func effectSpecs() []spriteSpec {
	return []spriteSpec{
		{
			id:            IDFxExplosionSmall,
			frameDuration: 6,
			frames: [][]string{
				{
					"..O..",
					".ORO.",
					"ORYRO",
					".ORO.",
					"..O..",
				},
				{
					".O.O.",
					"O.R.O",
					".RYR.",
					"O.R.O",
					".O.O.",
				},
				{
					"O...O",
					".R.R.",
					"..Y..",
					".R.R.",
					"O...O",
				},
				{
					".....",
					".R.R.",
					"..R..",
					".R.R.",
					".....",
				},
			},
		},
		{
			id:            IDFxExplosionLarge,
			frameDuration: 6,
			frames: [][]string{
				{
					"....O....",
					"..OOROO..",
					".ORRYYR.O",
					"ORYYYYYRO",
					".ORRYYR.O",
					"..OOROO..",
					"....O....",
				},
				{
					"..O...O..",
					".ORR.RRO.",
					"ORYYYYYRO",
					".RYYYYYR.",
					"ORYYYYYRO",
					".ORR.RRO.",
					"..O...O..",
				},
				{
					".O.....O.",
					"O..R.R..O",
					".R.....R.",
					"..R.Y.R..",
					".R.....R.",
					"O..R.R..O",
					".O.....O.",
				},
				{
					"O.......O",
					"..R...R..",
					"...R.R...",
					"....Y....",
					"...R.R...",
					"..R...R..",
					"O.......O",
				},
				{
					".........",
					".R.....R.",
					"...R.R...",
					"....R....",
					"...R.R...",
					".R.....R.",
					".........",
				},
				{
					".........",
					"....R....",
					".........",
					"...R.R...",
					".........",
					"....R....",
					".........",
				},
			},
		},
		{
			id:            IDFxSparkle,
			frameDuration: 12,
			frames: [][]string{
				{
					"..C..",
					"..C..",
					"CCCCC",
					"..C..",
					"..C..",
				},
				{
					"C...C",
					".C.C.",
					"..C..",
					".C.C.",
					"C...C",
				},
			},
		},
	}
}

func recolorSpecs(frames [][]string, from, to rune) [][]string {
	out := make([][]string, 0, len(frames))
	for _, frame := range frames {
		rows := make([]string, 0, len(frame))
		for _, row := range frame {
			r := []rune(row)
			for i, ch := range r {
				if ch == from {
					r[i] = to
				}
			}
			rows = append(rows, string(r))
		}
		out = append(out, rows)
	}
	return out
}
