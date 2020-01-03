package collision

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type BoundingCircle struct {
	pixel.Circle
}

func NewBoundingCircle(x, y, r float64) *BoundingCircle {
	return &BoundingCircle{
		pixel.Circle{
			Center: pixel.V(x, y),
			Radius: r,
		},
	}
}

func (bc *BoundingCircle) Collides(other Shaper) bool {
	switch t := other.Shape().(type) {
	case *BoundingBox:
		return bc.IntersectRect(t.Rect) != pixel.ZV
	case *BoundingCircle:
		return bc.Intersect(t.Circle).Radius != 0
	}
	return false
}

func (bc *BoundingCircle) Shape() Shape {
	return bc
}

// TODO
func (bb *BoundingCircle) Align(pos pixel.Vec) {
	return
}

// TODO
func (bc *BoundingCircle) Resolve(delta pixel.Vec, others ...Shaper) Solution {
	sol := Solution{}

	return sol
}

// This code is wrong, get the one from bb
func (bc *BoundingCircle) resolveAgainstBoundingBox(other *BoundingBox, distance pixel.Vec) Solution {
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

func (bc *BoundingCircle) Draw(color *color.RGBA, imd *imdraw.IMDraw, target pixel.Target) {
	imd.Color = *color
	imd.Push(bc.Center)
	imd.Circle(bc.Radius, 0)
}
