package animation

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/faiface/pixel"
	"github.com/svera/quarter"
	"github.com/svera/quarter/collision"
)

// BoundedAnimFile is the struct of the JSON file used to store an animated sprite with attached bounding boxes
type BoundedAnimFile struct {
	Version string
	Sheet   string
	Anims   []struct {
		AnimData
		Boxes []collision.BoundingBox
	}
}

// Character is a wrapper around both AnimSprite and BoundingBox that associates
// both to make "solid" animated sprites
type Character struct {
	AnimSprite
	BoundingBoxes map[int][]collision.BoundingBox
}

// LoadCharacter loads a character data from reader
func LoadCharacter(r io.Reader, pos pixel.Vec) (*Character, error) {
	chtr := &Character{
		BoundingBoxes: make(map[int][]collision.BoundingBox),
	}
	data := &BoundedAnimFile{}
	err := json.NewDecoder(r).Decode(data)
	if err != nil {
		return chtr, err
	}
	if data.Version != "1" {
		return chtr, fmt.Errorf("Version not supported")
	}
	pic, err := quarter.LoadPicture(data.Sheet)
	if err != nil {
		return chtr, err
	}
	chtr.AnimSprite = *NewAnimSprite(pos, len(data.Anims))
	for i, an := range data.Anims {
		chtr.AddAnim(i, pic, an.YOffset, an.Width, an.Height, an.Frames, an.Duration, an.Cycle)
		for _, bb := range an.Boxes {
			chtr.BoundingBoxes[i] = append(chtr.BoundingBoxes[i], bb)
		}
	}
	return chtr, nil
}

// BoundingBox returns the character bounding box information updated to its current position
func (c *Character) BoundingBox() *collision.BoundingBox {
	bb := c.BoundingBoxes[c.currentAnimID][c.currentFrameNumber]
	bb.Rect = bb.Moved(c.Position)
	return &bb
}

func (c *Character) InBoundsX(delta pixel.Vec, limits pixel.Rect) bool {
	return c.BoundingBox().Min.X+delta.X > limits.Min.X &&
		c.BoundingBox().Max.X+delta.X < limits.Max.X
}

func (c *Character) InBoundsY(delta pixel.Vec, limits pixel.Rect) bool {
	return c.BoundingBox().Min.Y+delta.Y > limits.Min.Y &&
		c.BoundingBox().Max.Y+delta.Y < limits.Max.Y
}
