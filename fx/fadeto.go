package fx

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Fade struct {
	Color    pixel.RGBA
	elapsed  float64
	total    float64
	duration float64
	step     float64
}

func NewFade(color pixel.RGBA, duration float64) *Fade {
	color.A = 0
	return &Fade{
		Color:    color,
		step:     duration / 1000,
		duration: duration,
	}
}

func (t *Fade) To(win *pixelgl.Window, imd *imdraw.IMDraw, dt float64) bool {
	t.elapsed += dt
	t.total += dt
	if t.total > t.duration {
		return true
	}
	if t.elapsed > t.step {
		t.elapsed = 0
		t.Color.A += dt / t.duration
	}
	imd.Color = t.Color
	imd.Push(
		pixel.ZV,
		pixel.V(win.Bounds().W(), win.Bounds().H()),
	)
	imd.Rectangle(0)
	imd.Draw(win)
	return false
}
