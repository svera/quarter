package scene

import "github.com/faiface/pixel/pixelgl"

// Scene provides a way for the user to get from one scene to the next
// Each scene is a distinct state of the game that displays different information
type Scene interface {
	// Loop executes scene logic and returns name of next scene to be run
	Loop(w *pixelgl.Window, dt float64) (string, error)
}
