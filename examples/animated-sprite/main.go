package main

import (
	"bytes"
	"fmt"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/svera/quarter"
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
	width  = 256
	height = 192
	// To achieve this pixelated retro look, we use internally a 256x192 canvas and scale it up to 1024x768 (4x) when showing on screen
	zoom = 4
)

const (
	heroWidth  = 19
	heroHeight = 22
	impulse    = 64
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Animation demo",
		Bounds: pixel.R(0, 0, width*zoom, height*zoom),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	r, err := os.Open("hero.json")
	if err != nil {
		panic(err)
	}

	face, err := quarter.LoadTTF("assets/ARCADE_N.TTF", 21)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	txt := text.New(pixel.V(100, 500), atlas)
	txt.Color = color.Black
	fmt.Fprintln(txt, "Arcade")

	phys := physic.NewPhysics(
		physic.Params{
			MaxVelocityX: 50,
			Acceleration: 75,
			Gravity:      112,
		},
	)

	hero, err := animation.LoadCharacter(r, 64, 32)
	if err != nil {
		panic(err)
	}
	circle := collision.NewBoundingCircle(0, 40, 18)

	// Canvas origin of coordinates will be at its center
	canvas := pixelgl.NewCanvas(pixel.R(-width/2, -height/2, width/2, height/2))

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	imd := imdraw.New(nil)
	red := &color.RGBA{255, 0, 0, 16}
	green := &color.RGBA{0, 255, 0, 16}
	blue := &color.RGBA{0, 0, 255, 16}

	data, err := ioutil.ReadFile("levels.json")
	lvl, err := level.Load(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	lvl[0].SetDebug(imd, blue)

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		imd.Clear()

		readInput(hero, phys, win, dt)
		delta := phys.Displacement(dt)
		sol := hero.BoundingBox().Resolve(delta, lvl[0].Layers[0].Bounds...)

		updatePosition(hero, phys, sol, delta)
		updateAnim(hero, phys, sol)

		canvas.Clear(colornames.Skyblue)
		lvl[0].Draw(canvas)
		hero.Draw(canvas, dt)
		hero.BoundingBox().Draw(red, imd, canvas)
		circle.Draw(green, imd, canvas)
		imd.Draw(canvas)
		canvas.Draw(win, pixel.IM.Moved(win.Bounds().Center()).Scaled(win.Bounds().Center(), zoom))
		txt.Draw(win, pixel.IM)
		win.Update()

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}

func readInput(hero *animation.Character, phys *physic.Physics, win *pixelgl.Window, dt float64) {
	if win.JustPressed(pixelgl.KeyUp) && hero.GetCurrentAnim() != jumping && hero.GetCurrentAnim() != falling {
		phys.Jump(impulse)
	} else if win.Pressed(pixelgl.KeyLeft) {
		hero.Dir = physic.DirectionLeft
		phys.Accelerate(hero.Dir, dt)
	} else if win.Pressed(pixelgl.KeyRight) {
		hero.Dir = physic.DirectionRight
		phys.Accelerate(hero.Dir, dt)
	} else if !win.Pressed(pixelgl.KeyLeft) && !win.Pressed(pixelgl.KeyRight) && !phys.IsStopped() {
		phys.Decelerate(dt)
	}
}

func updatePosition(hero *animation.Character, phys *physic.Physics, sol collision.Solution, delta pixel.Vec) {
	if sol.CollisionAxis == collision.AxisX {
		phys.StopMovingX()
		hero.Position = hero.Position.Add(pixel.V(sol.Distance.X, delta.Y))
	} else if sol.CollisionAxis == collision.AxisY {
		phys.StopMovingY()
		hero.Position = hero.Position.Add(pixel.V(delta.X, sol.Distance.Y))
	} else if sol.CollisionAxis == collision.AxisBoth {
		if phys.IsMovingUp() {
			phys.StopMovingY()
		}
		phys.StopMovingX()
		hero.Position = hero.Position.Add(sol.Distance)
	} else {
		hero.Position = hero.Position.Add(delta)
	}
}

func updateAnim(hero *animation.Character, phys *physic.Physics, sol collision.Solution) {
	if phys.IsMovingUp() {
		hero.SetCurrentAnim(jumping)
	} else if phys.IsMovingDown() && sol.CollisionAxis != collision.AxisY {
		hero.SetCurrentAnim(falling)
	} else if !phys.IsStopped() {
		hero.SetCurrentAnim(running)
	} else {
		hero.SetCurrentAnim(idle)
	}
}

func main() {
	pixelgl.Run(run)
}
