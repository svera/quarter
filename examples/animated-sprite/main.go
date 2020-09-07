package main

import (
	"fmt"
	_ "image/png"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/svera/quarter/scene"
)

const (
	width  = 256
	height = 192
	// To achieve this pixelated retro look, we use internally a 256x192 canvas and scale it up to 1024x768 (4x) when showing on screen
	zoom = 4
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

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	// Canvas origin of coordinates will be at its center
	canvas := pixelgl.NewCanvas(pixel.R(-width/2, -height/2, width/2, height/2))
	imd := imdraw.New(nil)

	scenes := map[string]scene.Scene{
		"attract": NewAttract(imd),
		"game":    NewGame(canvas, imd),
	}
	last := time.Now()

	currentScene := "attract"
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		if currentScene, err = scenes[currentScene].Loop(win, dt); err != nil {
			panic(err)
		}
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

func main() {
	pixelgl.Run(run)
}
