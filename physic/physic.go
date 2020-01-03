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

// Params is a set of values used to calculate velocuty and displacement in both axis
type Params struct {
	MaxVelocityX  float64
	MaxVelocityY  float64
	AccelerationX float64
	AccelerationY float64
	Gravity       float64
}

// Physics controls movement of an element
type Physics struct {
	// Speed of movement of the item in the X (horizontal) axis
	// negative values will move item to the left, positive ones to the right
	velocityX float64

	// Speed of movement of the item in the Y (vertical) axis
	// negative values will move item down, positive ones up
	velocityY float64

	Params
}

// NewPhysics returns a new instance of physics
func NewPhysics(params Params) *Physics {
	return &Physics{
		Params: params,
	}
}

// AccelerateX increases the speed of the item
func (p *Physics) AccelerateX(dir float64, dt float64) {
	p.velocityX += dir * p.AccelerationX * dt
	if math.Abs(p.velocityX) > math.Abs(p.MaxVelocityX) {
		p.velocityX = p.MaxVelocityX * dir
	}
}

// AccelerateY increases the speed of the item
func (p *Physics) AccelerateY(dir float64, dt float64) {
	p.velocityY += dir * p.AccelerationY * dt
	if math.Abs(p.velocityY) > math.Abs(p.MaxVelocityY) {
		p.velocityY = p.MaxVelocityY * dir
	}
}

// DecelerateX slow down velocity of the item in the X axis
func (p *Physics) DecelerateX(dt float64) {
	// We pick the minimum between velocity and acceleration
	// to avoid decelerating too much and never stopping completely the object
	val := math.Min(math.Abs(p.velocityX), p.AccelerationX*dt)
	if p.velocityX > 0 {
		p.velocityX -= val
	}

	if p.velocityX < 0 {
		p.velocityX += val
	}
}

// DecelerateY slow down velocity of the item in the Y axis
func (p *Physics) DecelerateY(dt float64) {
	// We pick the minimum between velocity and acceleration
	// to avoid decelerating too much and never stopping completely the object
	val := math.Min(math.Abs(p.velocityY), p.AccelerationY*dt)
	if p.velocityY > 0 {
		p.velocityY -= val
	}

	if p.velocityY < 0 {
		p.velocityY += val
	}
}

// SetVelocityX sets the horizontal speed of the item
func (p *Physics) SetVelocityX(value float64) {
	if math.Abs(value) < math.Abs(p.MaxVelocityX) {
		p.velocityX = value
	} else {
		p.velocityX = p.MaxVelocityX
	}
}

// SetVelocityY sets the vertical speed of the item
func (p *Physics) SetVelocityY(value float64) {
	if math.Abs(value) < math.Abs(p.MaxVelocityY) {
		p.velocityY = value
	} else {
		p.velocityY = p.MaxVelocityY
	}
}

// Displacement returns the movement an item must do both in X and Y axis
// after a dt time has passed
func (p *Physics) Displacement(dt float64) pixel.Vec {
	p.velocityY -= p.Gravity * dt
	return pixel.V(p.velocityX*dt, p.velocityY*dt)
}

// VelocityX returns object speed in X axis
func (p *Physics) VelocityX() float64 {
	return p.velocityX
}

// VelocityY returns object speed in Y axis
func (p *Physics) VelocityY() float64 {
	return p.velocityY
}
