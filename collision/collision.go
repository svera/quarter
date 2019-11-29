package collision

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Solution holds information about a collision if it happened, in which axis did and
// what's the maximum distance an object can move without colliding with the target object
type Solution struct {
	CollisionAxis int
	Object        Shaper
	Distance      pixel.Vec
}

// Possible collision axis values
const (
	AxisNone = iota
	AxisX    = 1
	AxisY    = 2
	AxisBoth = 3
)

type Shaper interface {
	Shape() Shape
}

type Shape interface {
	Collides(Shaper) bool
	Resolve(pixel.Vec, ...Shaper) Solution
	Draw(color *color.RGBA, imd *imdraw.IMDraw, target pixel.Target)
}