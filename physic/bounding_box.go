package physic

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// BoundingBox is a Rect with methods to resolve collisions
type BoundingBox struct {
	pixel.Rect
}

// NewBoundingBox returns a new BoundingBox instance. As Pixel's sprites,
// bounding box coordinates origin is located in its center.
func NewBoundingBox(x, y, w, h float64) *BoundingBox {
	return &BoundingBox{
		pixel.Rect{
			Min: pixel.V(x-(w/2), y-(h/2)),
			Max: pixel.V(x+(w/2), y+(h/2)),
		},
	}
}

func (bb *BoundingBox) Collides(other Shaper) bool {
	switch t := other.Shape().(type) {
	case *BoundingBox:
		return bb.Intersect(t.Rect) != pixel.ZR
	case *BoundingCircle:
		return bb.IntersectCircle(t.Circle) != pixel.ZV
	}
	return false
}

func (bb *BoundingBox) Shape() Shape {
	return bb
}

// Resolve checks if the bounding box will collide with another bounding box if it moves
// a certain delta
func (bb *BoundingBox) Resolve(delta pixel.Vec, others ...Shaper) CollisionSolution {
	sol := CollisionSolution{}

	for _, other := range others {
		switch t := other.Shape().(type) {
		case *BoundingBox:
			sol = bb.resolveAgainstBoundingBox(t, delta)
			if sol.CollisionAxis != CollisionNone {
				return sol
			}
		case *BoundingCircle:
			sol = bb.resolveAgainstBoundingCircle(t, delta)
			if sol.CollisionAxis != CollisionNone {
				return sol
			}
		}
	}
	return sol
}

func (bb *BoundingBox) resolveAgainstBoundingBox(other *BoundingBox, delta pixel.Vec) CollisionSolution {
	bbMoved := &BoundingBox{
		bb.Moved(delta),
	}

	intersectRect := bbMoved.Intersect(other.Rect)
	sol := CollisionSolution{
		Object:        other,
		CollisionAxis: CollisionNone,
	}

	if intersectRect == pixel.ZR {
		return sol
	}

	distance := bbMoved.Center().To(other.Center())

	if math.Abs(intersectRect.W()) <= math.Abs(intersectRect.H()) {
		if distance.X < 0 {
			sol.Distance.X = other.right() - bb.left()
		} else {
			sol.Distance.X = -(bb.right() - other.left())
		}
		sol.CollisionAxis = CollisionX
	} else {
		if distance.Y > 0 {
			sol.Distance.Y = other.bottom() - bb.top()
		} else {
			sol.Distance.Y = -(bb.bottom() - other.top())
		}
		sol.CollisionAxis = CollisionY
	}
	return sol
}

func (bb *BoundingBox) resolveAgainstBoundingCircle(other *BoundingCircle, delta pixel.Vec) CollisionSolution {
	bbMoved := &BoundingBox{
		bb.Moved(delta),
	}

	distance := bbMoved.Center().To(other.Center)

	d := other.Circle.IntersectRect(bbMoved.Rect)

	if d == pixel.ZV {
		return CollisionSolution{
			CollisionAxis: CollisionNone,
		}
	}

	sol := CollisionSolution{
		Object:        other,
		Distance:      d,
		CollisionAxis: CollisionBoth,
	}

	if math.Abs(distance.X) > math.Abs(distance.Y) && d == pixel.ZV {
		sol.CollisionAxis = CollisionX
	}

	if math.Abs(distance.X) < math.Abs(distance.Y) && d == pixel.ZV {
		sol.CollisionAxis = CollisionY
	}

	return sol
}

// left returns BoundBox's left side X coordinate
func (bb *BoundingBox) left() float64 {
	return bb.Min.X
}

// right returns BoundBox's right side X coordinate
func (bb *BoundingBox) right() float64 {
	return bb.Max.X
}

// top returns BoundBox's top side Y coordinate
func (bb *BoundingBox) top() float64 {
	return bb.Max.Y
}

// bottom returns BoundBox's bottom side Y coordinate
func (bb *BoundingBox) bottom() float64 {
	return bb.Min.Y
}

// Draw draws the bounding box surface on the passed target with the specified color for debugging purposes
func (bb *BoundingBox) Draw(color color.RGBA, imd *imdraw.IMDraw, target pixel.Target) {
	imd.Color = color
	imd.Push(
		pixel.V(bb.left(), bb.bottom()),
		pixel.V(bb.right(), bb.top()),
	)
	imd.Rectangle(0)
}
