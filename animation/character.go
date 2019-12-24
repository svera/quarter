package animation

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/faiface/pixel"
	"github.com/svera/quarter"
	"github.com/svera/quarter/collision"
)

// Returned errors
const (
	ErrorVersionNotSupported   = "Version \"%s\" not supported"
	ErrorNoAnims               = "File must have at least one animation declared, none found"
	ErrorShapeTypeNotSupported = "Shape type \"%s\" is not supported"
	ErrorShapeDataNotValid     = "Shape data of shape type \"%s\" is not valid"
)

// BoundedAnimFile is the struct of the JSON file used to store an animated sprite with attached bounding boxes
type BoundedAnimFile struct {
	Version string
	Sheet   string
	Anims   []struct {
		AnimData
		BoundingShapes []BoundingShape `json:"bounding_shapes"`
	}
}

// BoundingShape is a generic struct which can contain information about different shapes
type BoundingShape struct {
	Type       string
	Parameters json.RawMessage
}

// Character is a wrapper around both AnimSprite and BoundingBox that associates
// both to make "solid" animated sprites
type Character struct {
	AnimSprite
	BoundingShapes map[int][]collision.Shape
}

// LoadCharacter loads a character data from reader
func LoadCharacter(r io.Reader, pos pixel.Vec) (*Character, error) {
	chtr := &Character{
		BoundingShapes: make(map[int][]collision.Shape),
	}
	data := &BoundedAnimFile{}
	err := json.NewDecoder(r).Decode(data)

	if err != nil {
		return chtr, err
	}

	if data.Version != "1" {
		return chtr, fmt.Errorf(ErrorVersionNotSupported, data.Version)
	}

	pic, err := quarter.LoadPicture(data.Sheet)
	if err != nil {
		return chtr, err
	}

	if len(data.Anims) == 0 {
		return nil, fmt.Errorf(ErrorNoAnims)
	}

	chtr.AnimSprite = *NewAnimSprite(pos, len(data.Anims))
	for i, an := range data.Anims {
		chtr.AddAnim(i, pic, an.YOffset, an.Width, an.Height, an.Frames, an.Duration, an.Cycle)
		for _, shape := range an.BoundingShapes {
			switch shape.Type {
			case "box":
				bb := collision.BoundingBox{}
				err := json.Unmarshal(shape.Parameters, &bb)
				if err != nil {
					return nil, fmt.Errorf(ErrorShapeDataNotValid, shape.Type)
				}
				chtr.BoundingShapes[i] = append(chtr.BoundingShapes[i], &bb)
			default:
				return nil, fmt.Errorf(ErrorShapeTypeNotSupported, shape.Type)
			}
		}
	}
	return chtr, nil
}

// BoundingShape returns the character bounding box information updated to its current position
func (c *Character) BoundingShape() collision.Shape {
	bb := c.BoundingShapes[c.currentAnimID][c.currentFrameNumber]
	bb.Recenter(c.Position)
	return bb
}

/*
func (c *Character) InBoundsX(delta pixel.Vec, limits pixel.Rect) bool {
	return c.BoundingShape().Min.X+delta.X > limits.Min.X &&
		c.BoundingShape().Max.X+delta.X < limits.Max.X
}

func (c *Character) InBoundsY(delta pixel.Vec, limits pixel.Rect) bool {
	return c.BoundingShape().Min.Y+delta.Y > limits.Min.Y &&
		c.BoundingShape().Max.Y+delta.Y < limits.Max.Y
}
*/
