package main

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/svera/quarter"
	"github.com/svera/quarter/textfx"
)

type Attract struct {
	txt   *text.Text
	txtFx *textfx.TextFX
}

func NewAttract() *Attract {
	face, err := quarter.LoadTTF("assets/ARCADE_N.TTF", 21)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	a := Attract{
		txt:   text.New(pixel.V(100, 500), atlas),
		txtFx: textfx.NewTextFX(0.5),
	}
	a.txt.Color = color.White
	fmt.Fprintln(a.txt, "Arcade")
	return &a
}

func (a *Attract) Loop(w *pixelgl.Window, dt float64) string {
	w.Clear(color.Black)
	a.txtFx.DrawBlinking(a.txt, w, pixel.IM, dt)
	if w.JustPressed(pixelgl.KeySpace) {
		return "game"
	}
	return "attract"
}