package textfx

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

type TextFX struct {
	Frequency float64
	elapsed   float64
	show      bool
}

func NewTextFX(freq float64) *TextFX {
	return &TextFX{
		Frequency: freq,
		show:      true,
	}
}

func (t *TextFX) DrawBlinking(txt *text.Text, tgt pixel.Target, matrix pixel.Matrix, dt float64) {
	t.elapsed += dt
	if t.elapsed > t.Frequency {
		fmt.Println(t.show)
		t.show = !t.show
		t.elapsed = 0
	}
	if t.show {
		txt.Draw(tgt, matrix)
	}
}
