package bound

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Circle struct {
	pixel.Circle
}

func NewCircle(x, y, r float64) *Circle {
	return &Circle{
		pixel.Circle{
			Center: pixel.V(x, y),
			Radius: r,
		},
	}
}

func (bc *Circle) Collides(other Shaper) bool {
	switch t := other.Shape().(type) {
	case *Box:
		return bc.IntersectRect(t.Rect) != pixel.ZV
	case *Circle:
		return bc.Intersect(t.Circle).Radius != 0
	}
	return false
}

func (bc *Circle) Shape() Shape {
	return bc
}

// TODO
func (bb *Circle) Align(pos pixel.Vec) {
	return
}

// TODO
func (bc *Circle) Resolve(delta pixel.Vec, others ...Shaper) Solution {
	sol := Solution{}

	return sol
}

// This code is wrong, get the one from bb
func (bc *Circle) resolveAgainstBoundingBox(other *Box, distance pixel.Vec) Solution {
	sol := Solution{
		Object: other,
	}

	if math.Abs(distance.X) > math.Abs(distance.Y) {
		if distance.X < 0 {
			sol.Distance.X = other.right() - bc.Radius
		} else {
			sol.Distance.X = -(bc.Radius - other.left())
		}
		sol.CollisionAxis = AxisX
	} else {
		if distance.Y > 0 {
			sol.Distance.Y = other.top() - bc.Radius
		} else {
			sol.Distance.Y = -(bc.Radius - other.top())
		}
		sol.CollisionAxis = AxisY
	}
	return sol
}

func (bc *Circle) Draw(color *color.RGBA, imd *imdraw.IMDraw, target pixel.Target) {
	imd.Color = *color
	imd.Push(bc.Center)
	imd.Circle(bc.Radius, 0)
}
