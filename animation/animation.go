package animation

import (
	"fmt"

	"github.com/faiface/pixel"
)

// AnimFile defines the structure of a disk file containing information about animations
type AnimFile struct {
	Version string
	Sheet   string
	Anims   []struct {
		AnimData
	}
}

type AnimData struct {
	Frames   int
	Cycle    string
	Duration float64
	YOffset  float64 `json:"y_offset"`
	Width    float64
	Height   float64
}

// Animation types of cycles
const (
	SingleReverse = iota - 2
	CircularReverse
	Circular
	Single
)

// Returned errors
const (
	ErrorAnimationDoesNotExist = "Animation %d does not exist"
)

// To convert string values used in sprite definition file to integer values used internally
var animationCycle = map[string]int{"single_reverse": -2, "circular_reverse": -1, "circular": 0, "single": 1}

type sequence struct {
	frames       []*pixel.Sprite
	timePerFrame float64
	cycle        int
}

// Animation implements an animated sprite
type Animation struct {
	anims              []*sequence
	currentAnimID      int
	currentFrameNumber int
	elapsed            float64
	Position           pixel.Vec
	Dir                float64
}

// NewAnimation returns a new Sprite instance to be drawn at position x, y
func NewAnimation(pos pixel.Vec, numberAnims int) *Animation {
	return &Animation{
		anims:    make([]*sequence, numberAnims),
		Position: pos,
		Dir:      1,
	}
}

// AddAnim adds a new animation to the Sprite, identified with ID,
// whose frames are taken from pic from left to right, starting from X = 0
// duration defines how many seconds should it take for the animation to complete a cycle
func (a *Animation) AddAnim(idx int, pic pixel.Picture, yOffset, width, height float64, numberFrames int, duration float64, cycle string) {
	a.anims[idx] = &sequence{
		timePerFrame: duration / float64(numberFrames),
	}
	a.anims[idx].cycle = animationCycle[cycle]
	var x float64
	for i := 0; i < numberFrames; i++ {
		x = width * float64(i)
		a.anims[idx].frames = append(a.anims[idx].frames, pixel.NewSprite(pic, pixel.R(x, yOffset, x+width, yOffset+height)))
	}
}

// SetCurrentAnim defines which animation to play
func (a *Animation) SetCurrentAnim(ID int) error {
	if ID < 0 || ID > len(a.anims)-1 {
		return fmt.Errorf(ErrorAnimationDoesNotExist, ID)
	}
	if ID != a.currentAnimID {
		a.currentAnimID = ID
		a.currentFrameNumber = 0
		if a.anims[a.currentAnimID].cycle == SingleReverse || a.anims[a.currentAnimID].cycle == CircularReverse {
			a.currentFrameNumber = a.lastFrame()
		}
		a.elapsed = 0
	}
	return nil
}

// CurrentAnim returns the current animation index
func (a *Animation) CurrentAnim() int {
	return a.currentAnimID
}

// Draw draws Sprite current frame on target, and updates the former if needed
func (a *Animation) Draw(target pixel.Target, dt float64) {
	m := pixel.IM.ScaledXY(pixel.ZV, pixel.V(a.Dir, 1)).Moved(a.Position)
	a.anims[a.currentAnimID].frames[a.currentFrameNumber].Draw(target, m)
	a.elapsed += dt
	if idx := a.nextFrameIndex(dt); idx != a.currentFrameNumber {
		a.elapsed = 0
		a.currentFrameNumber = idx
	}
}

func (a *Animation) nextFrameIndex(dt float64) int {
	if a.elapsed <= a.anims[a.currentAnimID].timePerFrame {
		return a.currentFrameNumber
	}
	if a.anims[a.currentAnimID].cycle == CircularReverse {
		if a.currentFrameNumber == 0 {
			return len(a.anims[a.currentAnimID].frames) - 1
		}
		return a.currentFrameNumber - 1
	}

	if a.anims[a.currentAnimID].cycle == SingleReverse {
		if a.currentFrameNumber > 0 {
			return a.currentFrameNumber - 1
		}
		return a.currentFrameNumber
	}

	if a.isLastFrame(a.currentFrameNumber) && a.anims[a.currentAnimID].cycle == Circular {
		return 0
	}
	if !a.isLastFrame(a.currentFrameNumber) {
		return a.currentFrameNumber + 1
	}
	return a.currentFrameNumber
}

func (a *Animation) lastFrame() int {
	return len(a.anims[a.currentAnimID].frames) - 1
}

func (a *Animation) isLastFrame(number int) bool {
	return len(a.anims[a.currentAnimID].frames)-1 == number
}
