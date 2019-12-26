package main

import (
	"bytes"
	"image/color"
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/svera/quarter/animation"
	"github.com/svera/quarter/collision"
	"github.com/svera/quarter/level"
	"github.com/svera/quarter/physic"
	"golang.org/x/image/colornames"
)

// Animation status
const (
	idle = iota
	running
	jumping
	falling
)

const (
	impulse = 64
)

type Game struct {
	phys   *physic.Physics
	hero   *animation.BoundedAnimation
	circle *collision.BoundingCircle
	canvas *pixelgl.Canvas
	imd    *imdraw.IMDraw
	levels []level.Level
	red    *color.RGBA
	green  *color.RGBA
	blue   *color.RGBA
}

func NewGame(width, height float64) *Game {
	r, err := os.Open("hero.json")
	if err != nil {
		panic(err)
	}

	hero, err := animation.LoadBoundedAnimation(r, pixel.V(64, 32))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile("levels.json")
	lvl, err := level.Load(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	g := Game{
		phys: physic.NewPhysics(
			physic.Params{
				MaxVelocityX: 50,
				Acceleration: 75,
				Gravity:      112,
			},
		),
		hero:   hero,
		circle: collision.NewBoundingCircle(0, 40, 18),
		// Canvas origin of coordinates will be at its center
		canvas: pixelgl.NewCanvas(pixel.R(-width/2, -height/2, width/2, height/2)),
		imd:    imdraw.New(nil),
		levels: lvl,
	}

	// Canvas origin of coordinates will be at its center
	g.red = &color.RGBA{255, 0, 0, 16}
	g.green = &color.RGBA{0, 255, 0, 16}
	g.blue = &color.RGBA{0, 0, 255, 16}

	g.levels[0].SetDebug(g.imd, g.blue)
	return &g
}

func (g *Game) Loop(win *pixelgl.Window, dt float64) string {
	g.imd.Clear()

	g.readInput(win, dt)
	delta := g.phys.Displacement(dt)
	sol := g.hero.BoundingShape().Resolve(delta, g.levels[0].Layers[0].Bounds...)

	/*
		if !g.hero.InBoundsX(delta, g.levels[0].Limits) {
			delta.X = 0
			g.phys.StopMovingX()
		}
	*/
	g.updatePosition(sol, delta)
	g.updateAnim(sol)

	g.canvas.Clear(colornames.Skyblue)
	g.levels[0].Draw(g.canvas)
	g.hero.Draw(g.canvas, dt)
	g.hero.BoundingShape().Draw(g.red, g.imd, g.canvas)
	g.circle.Draw(g.green, g.imd, g.canvas)
	g.imd.Draw(g.canvas)
	g.canvas.Draw(win, pixel.IM.Moved(win.Bounds().Center()).Scaled(win.Bounds().Center(), zoom))
	return "game"
}

func (g *Game) readInput(win *pixelgl.Window, dt float64) {
	if win.JustPressed(pixelgl.KeyUp) && g.hero.GetCurrentAnim() != jumping && g.hero.GetCurrentAnim() != falling {
		g.phys.Jump(impulse)
	} else if win.Pressed(pixelgl.KeyLeft) {
		g.hero.Dir = physic.DirectionLeft
		g.phys.Accelerate(g.hero.Dir, dt)
	} else if win.Pressed(pixelgl.KeyRight) {
		g.hero.Dir = physic.DirectionRight
		g.phys.Accelerate(g.hero.Dir, dt)
	} else if !win.Pressed(pixelgl.KeyLeft) && !win.Pressed(pixelgl.KeyRight) && !g.phys.IsStopped() {
		g.phys.Decelerate(dt)
	}
}

func (g *Game) updatePosition(sol collision.Solution, delta pixel.Vec) {
	if sol.CollisionAxis == collision.AxisX {
		g.phys.StopMovingX()
		g.hero.Position = g.hero.Position.Add(pixel.V(sol.Distance.X, delta.Y))
	} else if sol.CollisionAxis == collision.AxisY {
		g.phys.StopMovingY()
		g.hero.Position = g.hero.Position.Add(pixel.V(delta.X, sol.Distance.Y))
	} else if sol.CollisionAxis == collision.AxisBoth {
		if g.phys.IsMovingUp() {
			g.phys.StopMovingY()
		}
		g.phys.StopMovingX()
		g.hero.Position = g.hero.Position.Add(sol.Distance)
	} else {
		g.hero.Position = g.hero.Position.Add(delta)
	}
}

func (g *Game) updateAnim(sol collision.Solution) {
	if g.phys.IsMovingUp() {
		g.hero.SetCurrentAnim(jumping)
	} else if g.phys.IsMovingDown() && sol.CollisionAxis != collision.AxisY {
		g.hero.SetCurrentAnim(falling)
	} else if !g.phys.IsStopped() {
		g.hero.SetCurrentAnim(running)
	} else {
		g.hero.SetCurrentAnim(idle)
	}
}
