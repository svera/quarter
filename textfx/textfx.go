package textfx

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

// Blinking implements text writing on screen with applied effects
type Blinking struct {
	Frequency float64
	elapsed   float64
	show      bool
}

// NewBlinking returns a new TextFX instance
func NewBlinking(freq float64) *Blinking {
	return &Blinking{
		Frequency: freq,
		show:      true,
	}
}

// Blinking draws passed text on screen, showing it on and off according to the specified frequency
func (t *Blinking) Blinking(txt *text.Text, tgt pixel.Target, matrix pixel.Matrix, dt float64) {
	t.elapsed += dt
	if t.elapsed > t.Frequency {
		t.show = !t.show
		t.elapsed = 0
	}
	if t.show {
		txt.Draw(tgt, matrix)
	}
}
