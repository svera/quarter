package physic

import (
	"math"

	"github.com/faiface/pixel"
)

// Movement direction for the item
const (
	DirectionLeft  = -1
	DirectionRight = 1
	DirectionUp    = 1
	DirectionDown  = -1
)

// Movement axis
const (
	AxisX = iota
	AxisY
)

// Params is a set of values used to calculate velocuty and displacement in both axis
type Params struct {
	MaxVelocity  [2]float64
	Acceleration [2]float64
	Gravity      float64
}

// Physics controls movement of an element
type Physics struct {
	// Speed of movement of the item in the X (horizontal) axis
	// negative values will move item to the left or down, positive ones to the right or up,
	// depending on the axys
	velocity [2]float64

	Params
}

// NewPhysics returns a new instance of physics
func NewPhysics(params Params) *Physics {
	return &Physics{
		Params: params,
	}
}

// Accelerate increases the speed of the item
func (p *Physics) Accelerate(axis int, dir float64, dt float64) {
	p.velocity[axis] += dir * p.Acceleration[axis] * dt
	if math.Abs(p.velocity[axis]) > math.Abs(p.MaxVelocity[axis]) {
		p.velocity[axis] = p.MaxVelocity[axis] * dir
	}
}

// Decelerate slow down velocity of the item in the passed axis
func (p *Physics) Decelerate(axis int, dt float64) {
	// We pick the minimum between velocity and acceleration
	// to avoid decelerating too much and never stopping completely the object
	val := math.Min(math.Abs(p.velocity[axis]), p.Acceleration[axis]*dt)
	if p.velocity[axis] > 0 {
		p.velocity[axis] -= val
	}

	if p.velocity[axis] < 0 {
		p.velocity[axis] += val
	}
}

// SetVelocity sets the speed of the item on the passed axis
func (p *Physics) SetVelocity(axis int, value float64) {
	if math.Abs(value) < math.Abs(p.MaxVelocity[axis]) {
		p.velocity[axis] = value
	} else {
		p.velocity[axis] = p.MaxVelocity[axis]
	}
}

// Displacement returns the movement an item must do both in X and Y axis
// after a dt time has passed
func (p *Physics) Displacement(dt float64) pixel.Vec {
	p.velocity[AxisY] -= p.Gravity * dt
	return pixel.V(p.velocity[AxisX]*dt, p.velocity[AxisY]*dt)
}

// Velocity returns object speed in the passed axis
func (p *Physics) Velocity(axis int) float64 {
	return p.velocity[axis]
}
