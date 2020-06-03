package main

import (
	"bytes"
	"image/color"
	"io"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/svera/quarter/animation"
	"github.com/svera/quarter/bound"
	"github.com/svera/quarter/physic"
)

const (
	impulse = 64
)

type Hero struct {
	*physic.Physics
	*animation.Animation
	heroBounds map[string][]bound.Shaper
}

func NewHero(dataFile string) (*Hero, error) {
	r, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	anim, err := animation.Deserialize(tee, pixel.V(64, 32))
	if err != nil {
		panic(err)
	}
	anim.SetCurrentAnim("idle")
	bounds, err := bound.Deserialize(&buf)
	if err != nil {
		return nil, err
	}

	return &Hero{
		physic.NewPhysics(
			physic.Params{
				MaxVelocity:  [2]float64{50, 100},
				Acceleration: [2]float64{75, 0},
				Gravity:      112,
			},
		),
		anim,
		bounds,
	}, nil
}

func (h *Hero) Jump(dt float64) {
	if h.CurrentAnim() != "jumping" && h.CurrentAnim() != "falling" {
		h.SetVelocity(physic.AxisY, impulse)
	}
}

func (h *Hero) Left(dt float64) {
	h.Dir = physic.DirectionLeft
	h.Accelerate(physic.AxisX, h.Dir, dt)
}

func (h *Hero) Right(dt float64) {
	h.Dir = physic.DirectionRight
	h.Accelerate(physic.AxisX, h.Dir, dt)
}

func (h *Hero) Draw(target *pixelgl.Canvas, debug *color.RGBA, imd *imdraw.IMDraw, dt float64) {
	h.Animation.Draw(target, dt)
	h.boundingShape().Draw(debug, imd, target)
}

func (h *Hero) updatePosition(sol bound.Solution, delta pixel.Vec) {
	if sol.CollisionAxis == bound.AxisX {
		h.SetVelocity(physic.AxisX, 0)
		h.Position = h.Position.Add(pixel.V(sol.Distance.X, delta.Y))
	} else if sol.CollisionAxis == bound.AxisY {
		h.SetVelocity(physic.AxisY, 0)
		h.Position = h.Position.Add(pixel.V(delta.X, sol.Distance.Y))
	} else if sol.CollisionAxis == bound.AxisBoth {
		h.SetVelocity(physic.AxisY, 0)
		h.SetVelocity(physic.AxisX, 0)
		h.Position = h.Position.Add(sol.Distance)
	} else {
		h.Position = h.Position.Add(delta)
	}
	h.updateAnim(sol)
}

func (h *Hero) boundingShape() bound.Shape {
	bb := h.heroBounds[h.CurrentAnim()][h.CurrentFrameNumber()]
	bb.Shape().Align(h.Position)
	return bb.Shape()
}

func (h *Hero) updateAnim(sol bound.Solution) {
	if h.Velocity(physic.AxisY) > 0 {
		h.SetCurrentAnim("jumping")
	} else if h.Velocity(physic.AxisY) < 0 && sol.CollisionAxis != bound.AxisY {
		h.SetCurrentAnim("falling")
	} else if h.Velocity(physic.AxisX) != 0 {
		h.SetCurrentAnim("running")
	} else {
		h.SetCurrentAnim("idle")
	}
}
