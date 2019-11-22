package animation

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/svera/quarter"
	"github.com/svera/quarter/physic"
)

// BoundedAnimFile is the struct of the JSON file used to store an animated sprite with attached bounding boxes
type BoundedAnimFile struct {
	Version string
	Sheet   string
	Anims   []struct {
		AnimData
		Boxes []physic.BoundingBox
	}
}

type Character struct {
	AnimSprite
	BoundingBoxes map[int][]physic.BoundingBox
}

// LoadCharacter loads a character data from reader
func LoadCharacter(r io.Reader, x, y float64) (*Character, error) {
	chtr := &Character{
		BoundingBoxes: make(map[int][]physic.BoundingBox),
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
	chtr.AnimSprite = *NewAnimSprite(x, y, len(data.Anims))
	for i, an := range data.Anims {
		chtr.AddAnim(i, pic, an.YOffset, an.Width, an.Height, an.Frames, an.Duration, an.Cycle)
		for _, bb := range an.Boxes {
			chtr.BoundingBoxes[i] = append(chtr.BoundingBoxes[i], bb)
		}
	}
	return chtr, nil
}

// physic.BoundingBox returns the character bounding box information updated to its current position
func (c *Character) BoundingBox() *physic.BoundingBox {
	bb := c.BoundingBoxes[c.currentAnimID][c.currentFrameNumber]
	bb.Rect = bb.Moved(c.Position)
	return &bb
}
