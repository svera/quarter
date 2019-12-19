package textfx

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

// TextFX implements text writing on screen with applied effects
type TextFX struct {
	Frequency float64
	elapsed   float64
	show      bool
}

// NewTextFX returns a new TextFX instance
func NewTextFX(freq float64) *TextFX {
	return &TextFX{
		Frequency: freq,
		show:      true,
	}
}

// DrawBlinking draws passed text on screen, showing it on and off according to the specified frequency
func (t *TextFX) DrawBlinking(txt *text.Text, tgt pixel.Target, matrix pixel.Matrix, dt float64) {
	t.elapsed += dt
	if t.elapsed > t.Frequency {
		t.show = !t.show
		t.elapsed = 0
	}
	if t.show {
		txt.Draw(tgt, matrix)
	}
}
