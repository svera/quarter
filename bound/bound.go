package bound

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Returned errors
const (
	ErrorVersionNotSupported   = "Version \"%s\" not supported"
	ErrorShapeTypeNotSupported = "Shape type \"%s\" is not supported"
	ErrorShapeDataNotValid     = "Shape data of shape type \"%s\" is not valid"
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

// Shape is an interface that defines the minimum contract required for shapes that
// can be checked for collisions
type Shape interface {
	Collides(Shaper) bool
	Resolve(pixel.Vec, ...Shaper) Solution
	Draw(color *color.RGBA, imd *imdraw.IMDraw, target pixel.Target)
	Align(pos pixel.Vec)
}

// BoundFile is the struct of the JSON file used to store an animated sprite with attached bounding boxes
type BoundFile struct {
	Version string
	Bounds  map[string]struct {
		Shapes []struct {
			Type   string
			Values json.RawMessage
		}
	}
}

func Deserialize(r io.Reader) (map[string][]Shaper, error) {
	bounds := make(map[string][]Shaper)

	data := &BoundFile{}
	err := json.NewDecoder(r).Decode(data)

	if err != nil {
		return bounds, err
	}

	if data.Version != "1" {
		return bounds, fmt.Errorf(ErrorVersionNotSupported, data.Version)
	}

	for id, set := range data.Bounds {
		for _, shape := range set.Shapes {
			switch shape.Type {
			case "box":
				bb := Box{}
				err := json.Unmarshal(shape.Values, &bb)
				if err != nil {
					return nil, fmt.Errorf(ErrorShapeDataNotValid, shape.Type)
				}
				bounds[id] = append(bounds[id], &bb)
			default:
				return nil, fmt.Errorf(ErrorShapeTypeNotSupported, shape.Type)
			}
		}
	}

	return bounds, nil
}
