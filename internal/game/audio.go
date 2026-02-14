package game

import (
	"encoding/binary"
	"math"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const audioSampleRate = 44100

type audioEngine struct {
	shot      *voicePool
	explosion *voicePool
	bigBoom   *voicePool
	dive      *voicePool
}

type voicePool struct {
	players []*audio.Player
	next    int
}

func newAudioEngine() *audioEngine {
	ctx := audio.NewContext(audioSampleRate)

	return &audioEngine{
		shot:      newVoicePool(ctx, generateShotSound(), 4, 0.22),
		explosion: newVoicePool(ctx, generateSmallExplosionSound(), 4, 0.28),
		bigBoom:   newVoicePool(ctx, generateHugeExplosionSound(), 3, 0.45),
		dive:      newVoicePool(ctx, generateDiveSound(), 3, 0.26),
	}
}

func newVoicePool(ctx *audio.Context, pcm []byte, voices int, volume float64) *voicePool {
	v := &voicePool{players: make([]*audio.Player, 0, voices)}
	for i := 0; i < voices; i++ {
		p := audio.NewPlayerFromBytes(ctx, pcm)
		p.SetVolume(volume)
		v.players = append(v.players, p)
	}
	return v
}

func (v *voicePool) play() {
	if v == nil || len(v.players) == 0 {
		return
	}
	p := v.players[v.next]
	v.next = (v.next + 1) % len(v.players)
	_ = p.Rewind()
	p.Play()
}

func (a *audioEngine) playShot() {
	if a == nil {
		return
	}
	a.shot.play()
}

func (a *audioEngine) playSmallExplosion() {
	if a == nil {
		return
	}
	a.explosion.play()
}

func (a *audioEngine) playHugeExplosion() {
	if a == nil {
		return
	}
	a.bigBoom.play()
}

func (a *audioEngine) playDive() {
	if a == nil {
		return
	}
	a.dive.play()
}

func generateShotSound() []byte {
	const dur = 0.09
	return synthesizePCM(dur, func(t float64, _ *noiseSource) float64 {
		freq := 1500.0 - 1150.0*(t/dur)
		if freq < 260 {
			freq = 260
		}
		env := math.Exp(-24 * t)
		return (math.Sin(twoPi*freq*t) + 0.35*math.Sin(twoPi*freq*2.0*t)) * env * 0.55
	})
}

func generateSmallExplosionSound() []byte {
	const dur = 0.2
	var lp float64
	return synthesizePCM(dur, func(t float64, n *noiseSource) float64 {
		white := n.next()
		lp += (white - lp) * 0.22
		env := math.Exp(-10*t) * (1 - 0.2*t/dur)
		bass := math.Sin(twoPi*(170-90*(t/dur))*t) * math.Exp(-7*t)
		return (0.78*lp + 0.22*bass) * env * 0.95
	})
}

func generateHugeExplosionSound() []byte {
	const dur = 0.68
	var lp float64
	return synthesizePCM(dur, func(t float64, n *noiseSource) float64 {
		white := n.next()
		lp += (white - lp) * 0.08
		attack := math.Min(1, t*35)
		env := attack * math.Exp(-3.3*t)
		subFreq := 90.0 - 45.0*(t/dur)
		if subFreq < 34 {
			subFreq = 34
		}
		sub := math.Sin(twoPi*subFreq*t) * math.Exp(-2.2*t)
		crunch := math.Tanh(1.6 * lp)
		return (0.66*crunch + 0.34*sub) * env * 1.25
	})
}

func generateDiveSound() []byte {
	const dur = 0.45
	return synthesizePCM(dur, func(t float64, _ *noiseSource) float64 {
		norm := t / dur
		freq := 980 - 760*norm + 210*math.Sin(twoPi*3.6*t)
		if freq < 190 {
			freq = 190
		}
		env := math.Pow(math.Sin(math.Pi*norm), 0.8)
		wobble := 0.22 * math.Sin(twoPi*freq*1.9*t)
		return (math.Sin(twoPi*freq*t) + wobble) * env * 0.52
	})
}

const twoPi = math.Pi * 2

type noiseSource struct {
	state uint32
}

func (n *noiseSource) next() float64 {
	n.state = n.state*1664525 + 1013904223
	v := float64(n.state>>8) / float64(1<<24)
	return v*2 - 1
}

func synthesizePCM(duration float64, sampleFn func(t float64, n *noiseSource) float64) []byte {
	samples := int(float64(audioSampleRate) * duration)
	out := make([]byte, samples*4)
	noise := &noiseSource{state: 0x1234abcd}

	for i := 0; i < samples; i++ {
		t := float64(i) / audioSampleRate
		v := sampleFn(t, noise)
		if v > 1 {
			v = 1
		}
		if v < -1 {
			v = -1
		}
		s := int16(v * 32767)
		offset := i * 4
		binary.LittleEndian.PutUint16(out[offset:offset+2], uint16(s))
		binary.LittleEndian.PutUint16(out[offset+2:offset+4], uint16(s))
	}
	return out
}
