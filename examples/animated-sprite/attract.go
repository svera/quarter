package main

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/svera/quarter"
	"github.com/svera/quarter/fx"
)

type Attract struct {
	txt   *text.Text
	txtFx *fx.Blinking
	fade  *fx.Fade
	imd   *imdraw.IMDraw
	exit  bool
}

func NewAttract(imd *imdraw.IMDraw) *Attract {
	face, err := quarter.LoadTTF("assets/ARCADE_N.TTF", 21)
	if err != nil {
		panic(err)
	}

	centerX := float64((width * zoom) / 2)
	centerY := float64((height * zoom) / 2)
	atlas := text.NewAtlas(face, text.ASCII)
	black := pixel.RGB(0, 0, 0)
	a := Attract{
		txt:   text.New(pixel.V(centerX, centerY), atlas),
		txtFx: fx.NewBlinking(0.5),
		fade:  fx.NewFade(black.Mul(pixel.Alpha(0)), 1),
		imd:   imd,
		exit:  false,
	}
	a.txt.Color = color.White
	line := "Press a Key"
	a.txt.Dot.X -= a.txt.BoundsOf(line).W() / 2
	fmt.Fprintln(a.txt, line)
	return &a
}

func (a *Attract) Loop(w *pixelgl.Window, dt float64) (string, error) {
	a.imd.Clear()
	w.Clear(color.Black)
	a.txtFx.Draw(a.txt.Draw, w, pixel.IM, dt)
	if a.exit {
		if a.fade.To(w, a.imd, dt) {
			return "game", nil
		}
	}
	if w.JustPressed(pixelgl.KeySpace) {
		a.exit = true
	}
	return "attract", nil
}
