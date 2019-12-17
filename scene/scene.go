package scene

import "github.com/faiface/pixel/pixelgl"

type Scene interface {
	Loop(w *pixelgl.Window, dt float64) string // returns the name of the next scene
}
