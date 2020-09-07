package fx

import (
	"github.com/faiface/pixel"
)

// Blinking implements a flashing effect
type Blinking struct {
	Frequency float64
	elapsed   float64
	show      bool
}

// NewBlinking returns a new Blinking instance
func NewBlinking(freq float64) *Blinking {
	return &Blinking{
		Frequency: freq,
		show:      true,
	}
}

func (t *Blinking) Draw(draw func(pixel.Target, pixel.Matrix), tgt pixel.Target, matrix pixel.Matrix, dt float64) {
	t.elapsed += dt
	if t.elapsed > t.Frequency {
		t.show = !t.show
		t.elapsed = 0
	}
	if t.show {
		draw(tgt, matrix)
	}
}
