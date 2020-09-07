package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/svera/quarter/physic"
)

type Game struct {
	hero   *Hero
	level  *Level
	canvas *pixelgl.Canvas
	imd    *imdraw.IMDraw
}

func NewGame(canvas *pixelgl.Canvas, imd *imdraw.IMDraw) *Game {
	h, err := NewHero("hero.json")
	if err != nil {
		panic(err)
	}
	l, err := NewLevel("levels.json")
	if err != nil {
		panic(err)
	}

	g := Game{
		hero:  h,
		level: l,
		// Canvas origin of coordinates will be at its center
		canvas: canvas,
		imd:    imd,
	}

	return &g
}

func (g *Game) Loop(win *pixelgl.Window, dt float64) (string, error) {
	g.imd.Clear()
	g.readInput(win, dt)
	delta := g.hero.Displacement(dt)
	sol := g.hero.boundingShape().Resolve(delta, g.level.Bounds["level1-background"]...)

	g.hero.updatePosition(sol, delta)

	g.level.Draw(g.canvas, &color.RGBA{0, 0, 255, 16}, g.imd)
	g.hero.Draw(g.canvas, &color.RGBA{255, 0, 0, 16}, g.imd, dt)

	g.imd.Draw(g.canvas)
	g.canvas.Draw(win, pixel.IM.Moved(win.Bounds().Center()).Scaled(win.Bounds().Center(), zoom))
	return "game", nil
}

func (g *Game) readInput(win *pixelgl.Window, dt float64) {
	if win.JustPressed(pixelgl.KeyUp) {
		g.hero.Jump(dt)
	} else if win.Pressed(pixelgl.KeyLeft) {
		g.hero.Left(dt)
	} else if win.Pressed(pixelgl.KeyRight) {
		g.hero.Right(dt)
	} else if !win.Pressed(pixelgl.KeyLeft) && !win.Pressed(pixelgl.KeyRight) && g.hero.Velocity(physic.AxisX) != 0 {
		g.hero.Decelerate(physic.AxisX, dt)
	}
}
