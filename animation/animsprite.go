package animation

import (
	"fmt"

	"github.com/faiface/pixel"
)

type AnimFile struct {
	Version string
	Sheet   string
	Anims   []struct {
		AnimData
	}
}

type AnimData struct {
	Frames   float64
	Cycle    string
	Duration float64
	YOffset  float64 `json:"y_offset"`
	Width    float64
	Height   float64
}

const (
	Reverse = iota - 1
	Circular
	Single
)

// To convert string values used in sprite definition file to integer values
var animationCycle = map[string]int{"reverse": -1, "circular": 0, "single": 1}

type animation struct {
	frames       []*pixel.Sprite
	timePerFrame float64
	cycle        int
}

// AnimSprite implements an animated sprite
type AnimSprite struct {
	anims              []*animation
	currentAnimID      int
	currentFrameNumber int
	elapsed            float64
	Position           pixel.Vec
	Dir                float64
}

// NewAnimSprite returns a new AnimSprite instance to be drawn at position x, y
func NewAnimSprite(x, y float64, numberAnims int) *AnimSprite {
	return &AnimSprite{
		anims:    make([]*animation, numberAnims),
		Position: pixel.V(x, y),
		Dir:      1,
	}
}

// AddAnim adds a new animation to the AnimSprite, identified with ID,
// whose frames are taken from pic from left to right, starting from X = 0
// duration defines how many seconds should it take for the animation to complete a cycle
func (a *AnimSprite) AddAnim(idx int, pic pixel.Picture, yOffset, width, height, numberFrames, duration float64, cycle string) {
	a.anims[idx] = &animation{
		timePerFrame: duration / numberFrames,
	}
	a.anims[idx].cycle = animationCycle[cycle]
	var x float64
	for i := 0.0; i < numberFrames; i++ {
		x = width * i
		a.anims[idx].frames = append(a.anims[idx].frames, pixel.NewSprite(pic, pixel.R(x, yOffset, x+width, yOffset+height)))
	}
}

// SetCurrentAnim defines which animation to play
func (a *AnimSprite) SetCurrentAnim(ID int) error {
	if ID < 0 || ID > len(a.anims)-1 {
		return fmt.Errorf("Animation does not exist")
	}
	if ID != a.currentAnimID {
		a.currentAnimID = ID
		a.currentFrameNumber = 0
		if a.anims[a.currentAnimID].cycle == Reverse {
			a.currentFrameNumber = a.lastFrame()
		}
		a.elapsed = 0
	}
	return nil
}

func (a *AnimSprite) GetCurrentAnim() int {
	return a.currentAnimID
}

// Draw draws AnimSprite current frame on target, and updates the former if needed
func (a *AnimSprite) Draw(target pixel.Target, dt float64) {
	m := pixel.IM.ScaledXY(pixel.ZV, pixel.V(a.Dir, 1)).Moved(a.Position)
	a.anims[a.currentAnimID].frames[a.currentFrameNumber].Draw(target, m)
	a.elapsed += dt
	if idx := a.nextFrameIndex(dt); idx != a.currentFrameNumber {
		a.elapsed = 0
		a.currentFrameNumber = idx
	}
}

func (a *AnimSprite) nextFrameIndex(dt float64) int {
	if a.elapsed <= a.anims[a.currentAnimID].timePerFrame {
		return a.currentFrameNumber
	}
	if a.currentFrameNumber != 0 && a.anims[a.currentAnimID].cycle == Reverse {
		return a.currentFrameNumber - 1
	}
	if a.isLastFrame(a.currentFrameNumber) && a.anims[a.currentAnimID].cycle == Circular {
		return 0
	}
	if !a.isLastFrame(a.currentFrameNumber) {
		return a.currentFrameNumber + 1
	}
	return a.currentFrameNumber
}

func (a *AnimSprite) lastFrame() int {
	return len(a.anims[a.currentAnimID].frames) - 1
}

func (a *AnimSprite) isLastFrame(number int) bool {
	return len(a.anims[a.currentAnimID].frames)-1 == number
}
