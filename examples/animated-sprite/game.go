package main

import (
	"bytes"
	"image/color"
	"io"
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

const (
	impulse = 64
)

type Game struct {
	phys        *physic.Physics
	hero        *animation.Animation
	heroBounds  map[string][]collision.Shaper
	levelBounds map[string][]collision.Shaper
	circle      *collision.BoundingCircle
	canvas      *pixelgl.Canvas
	imd         *imdraw.IMDraw
	levels      map[string]level.Level
	debug       *color.RGBA
}

func NewGame(width, height float64) *Game {
	r, err := os.Open("hero.json")
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	hero, err := animation.Deserialize(tee, pixel.V(64, 32))
	if err != nil {
		panic(err)
	}
	hero.SetCurrentAnim("idle")
	heroBounds, err := collision.LoadBounds(&buf)
	if err != nil {
		panic(err)
	}

	data, err := os.Open("levels.json")
	if err != nil {
		panic(err)
	}
	tee = io.TeeReader(data, &buf)
	lvl, err := level.Deserialize(tee)
	if err != nil {
		panic(err)
	}
	levelBounds, err := collision.LoadBounds(&buf)
	if err != nil {
		panic(err)
	}

	g := Game{
		phys: physic.NewPhysics(
			physic.Params{
				MaxVelocity:  [2]float64{50, 100},
				Acceleration: [2]float64{75, 0},
				Gravity:      112,
			},
		),
		hero:        hero,
		heroBounds:  heroBounds,
		levelBounds: levelBounds,
		circle:      collision.NewBoundingCircle(0, 40, 18),
		// Canvas origin of coordinates will be at its center
		canvas: pixelgl.NewCanvas(pixel.R(-width/2, -height/2, width/2, height/2)),
		imd:    imdraw.New(nil),
		levels: lvl,
	}

	g.debug = &color.RGBA{0, 0, 255, 16}
	return &g
}

func (g *Game) Loop(win *pixelgl.Window, dt float64) (string, error) {
	g.imd.Clear()

	g.readInput(win, dt)
	delta := g.phys.Displacement(dt)
	sol := g.BoundingShape().Resolve(delta, g.levelBounds["level1-background"]...)

	g.updatePosition(sol, delta)
	g.updateAnim(sol)

	g.canvas.Clear(colornames.Skyblue)
	g.levels["level1"].Draw(g.canvas)
	if g.debug != nil {
		for _, b := range g.levelBounds["level1-background"] {
			b.Shape().Draw(g.debug, g.imd, g.canvas)
		}
	}

	g.hero.Draw(g.canvas, dt)
	g.BoundingShape().Draw(&color.RGBA{255, 0, 0, 16}, g.imd, g.canvas)
	g.circle.Draw(&color.RGBA{0, 255, 0, 16}, g.imd, g.canvas)
	g.imd.Draw(g.canvas)
	g.canvas.Draw(win, pixel.IM.Moved(win.Bounds().Center()).Scaled(win.Bounds().Center(), zoom))
	return "game", nil
}

func (g *Game) readInput(win *pixelgl.Window, dt float64) {
	if win.JustPressed(pixelgl.KeyUp) && g.hero.CurrentAnim() != "jumping" && g.hero.CurrentAnim() != "falling" {
		g.phys.SetVelocity(physic.AxisY, impulse)
	} else if win.Pressed(pixelgl.KeyLeft) {
		g.hero.Dir = physic.DirectionLeft
		g.phys.Accelerate(physic.AxisX, g.hero.Dir, dt)
	} else if win.Pressed(pixelgl.KeyRight) {
		g.hero.Dir = physic.DirectionRight
		g.phys.Accelerate(physic.AxisX, g.hero.Dir, dt)
	} else if !win.Pressed(pixelgl.KeyLeft) && !win.Pressed(pixelgl.KeyRight) && g.phys.Velocity(physic.AxisX) != 0 {
		g.phys.Decelerate(physic.AxisX, dt)
	}
}

func (g *Game) updatePosition(sol collision.Solution, delta pixel.Vec) {
	if sol.CollisionAxis == collision.AxisX {
		g.phys.SetVelocity(physic.AxisX, 0)
		g.hero.Position = g.hero.Position.Add(pixel.V(sol.Distance.X, delta.Y))
	} else if sol.CollisionAxis == collision.AxisY {
		g.phys.SetVelocity(physic.AxisY, 0)
		g.hero.Position = g.hero.Position.Add(pixel.V(delta.X, sol.Distance.Y))
	} else if sol.CollisionAxis == collision.AxisBoth {
		g.phys.SetVelocity(physic.AxisY, 0)
		g.phys.SetVelocity(physic.AxisX, 0)
		g.hero.Position = g.hero.Position.Add(sol.Distance)
	} else {
		g.hero.Position = g.hero.Position.Add(delta)
	}
}

func (g *Game) updateAnim(sol collision.Solution) {
	if g.phys.Velocity(physic.AxisY) > 0 {
		g.hero.SetCurrentAnim("jumping")
	} else if g.phys.Velocity(physic.AxisY) < 0 && sol.CollisionAxis != collision.AxisY {
		g.hero.SetCurrentAnim("falling")
	} else if g.phys.Velocity(physic.AxisX) != 0 {
		g.hero.SetCurrentAnim("running")
	} else {
		g.hero.SetCurrentAnim("idle")
	}
}

func (g *Game) BoundingShape() collision.Shape {
	bb := g.heroBounds[g.hero.CurrentAnim()][g.hero.CurrentFrameNumber()]
	bb.Shape().Align(g.hero.Position)
	return bb.Shape()
}
