package physic

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// CollisionSolution holds information about a collision if it happened, in which axis did and
// what's the maximum distance an object can move without colliding with the target object
type CollisionSolution struct {
	CollisionAxis int
	Object        Shaper
	Distance      pixel.Vec
}

// Possible collision axis values
const (
	CollisionNone = iota
	CollisionX    = 1
	CollisionY    = 2
	CollisionBoth = 3
)

type Shaper interface {
	Shape() Shape
}

type Shape interface {
	Collides(Shaper) bool
	Resolve(pixel.Vec, ...Shaper) CollisionSolution
	Draw(color color.RGBA, imd *imdraw.IMDraw, target pixel.Target)
}
