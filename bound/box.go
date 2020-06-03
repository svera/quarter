package bound

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Box is a Rect with methods to resolve collisions
type Box struct {
	pixel.Rect
}

// NewBox returns a new Box instance.
func NewBox(min pixel.Vec, max pixel.Vec) *Box {
	return &Box{
		pixel.Rect{
			Min: min,
			Max: max,
		},
	}
}

// Collides returns true if the bounding box collides with the passed shape
func (bb *Box) Collides(other Shaper) bool {
	switch t := other.Shape().(type) {
	case *Box:
		return bb.Intersects(t.Rect)
	case *Circle:
		return bb.IntersectCircle(t.Circle) != pixel.ZV
	}
	return false
}

// Shape returns the Box instance
func (bb *Box) Shape() Shape {
	return bb
}

func (bb *Box) Align(pos pixel.Vec) {
	bb.Rect = bb.Moved(pos.Sub(bb.Center()))
}

// Resolve checks if the bounding box will collide with another bounding box if it moves
// a certain delta
func (bb *Box) Resolve(delta pixel.Vec, others ...Shaper) Solution {
	sol := Solution{}

	for _, other := range others {
		switch t := other.Shape().(type) {
		case *Box:
			sol = bb.resolveAgainstBoundingBox(t, delta)
			if sol.CollisionAxis != AxisNone {
				return sol
			}
		case *Circle:
			sol = bb.resolveAgainstBoundingCircle(t, delta)
			if sol.CollisionAxis != AxisNone {
				return sol
			}
		}
	}
	return sol
}

func (bb *Box) resolveAgainstBoundingBox(other *Box, delta pixel.Vec) Solution {
	bbMoved := &Box{
		bb.Moved(delta),
	}

	sol := Solution{
		Object:        other,
		CollisionAxis: AxisNone,
	}

	if !bbMoved.Intersects(other.Rect) {
		return sol
	}

	intersectRect := bbMoved.Intersect(other.Rect)
	distance := bbMoved.Center().To(other.Center())

	if math.Abs(intersectRect.W()) <= math.Abs(intersectRect.H()) {
		if distance.X < 0 {
			sol.Distance.X = other.right() - bb.left()
		} else {
			sol.Distance.X = -(bb.right() - other.left())
		}
		sol.CollisionAxis = AxisX
	} else {
		if distance.Y > 0 {
			sol.Distance.Y = other.bottom() - bb.top()
		} else {
			sol.Distance.Y = -(bb.bottom() - other.top())
		}
		sol.CollisionAxis = AxisY
	}
	return sol
}

func (bb *Box) resolveAgainstBoundingCircle(other *Circle, delta pixel.Vec) Solution {
	bbMoved := &Box{
		bb.Moved(delta),
	}

	distance := bbMoved.Center().To(other.Center)

	d := other.Circle.IntersectRect(bbMoved.Rect)

	if d == pixel.ZV {
		return Solution{
			CollisionAxis: AxisNone,
		}
	}

	sol := Solution{
		Object:        other,
		Distance:      d,
		CollisionAxis: AxisBoth,
	}

	if math.Abs(distance.X) > math.Abs(distance.Y) && d == pixel.ZV {
		sol.CollisionAxis = AxisX
	}

	if math.Abs(distance.X) < math.Abs(distance.Y) && d == pixel.ZV {
		sol.CollisionAxis = AxisY
	}

	return sol
}

// left returns BoundBox's left side X coordinate
func (bb *Box) left() float64 {
	return bb.Min.X
}

// right returns BoundBox's right side X coordinate
func (bb *Box) right() float64 {
	return bb.Max.X
}

// top returns BoundBox's top side Y coordinate
func (bb *Box) top() float64 {
	return bb.Max.Y
}

// bottom returns BoundBox's bottom side Y coordinate
func (bb *Box) bottom() float64 {
	return bb.Min.Y
}

// Draw draws the bounding box surface on the passed target with the specified color for debugging purposes
func (bb *Box) Draw(color *color.RGBA, imd *imdraw.IMDraw, target pixel.Target) {
	imd.Color = *color
	imd.Push(
		pixel.V(bb.left(), bb.bottom()),
		pixel.V(bb.right(), bb.top()),
	)
	imd.Rectangle(0)
}
