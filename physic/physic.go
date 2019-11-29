package physic

import (
	"math"

	"github.com/faiface/pixel"
)

// Movement direction for the item
const (
	DirectionLeft  = -1
	DirectionRight = 1
)

type Params struct {
	MaxVelocityX float64
	Acceleration float64
	Gravity      float64
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

// Accelerate increases the speed of the item
func (p *Physics) Accelerate(dir float64, dt float64) {
	p.velocityX += dir * p.Acceleration * dt
	if math.Abs(p.velocityX) > math.Abs(p.MaxVelocityX) {
		p.velocityX = p.MaxVelocityX * dir
	}
	//fmt.Printf("velocity: %f\n", p.velocityX)
}

// Decelerate slow down velocity of the item
func (p *Physics) Decelerate(dt float64) {
	// We pick the minimum between velocity and acceleration
	// to avoid decelerating too much and never stopping completely the object
	val := math.Min(math.Abs(p.velocityX), p.Acceleration*dt)
	if p.velocityX > 0 {
		p.velocityX -= val
	}

	if p.velocityX < 0 {
		p.velocityX += val
	}
}

// Jump increases the vertical speed of the item
func (p *Physics) Jump(impulse float64) {
	p.velocityY = impulse
}

// Displacement returns the movement an item must do both in X and Y axis
// after a dt time has passed
func (p *Physics) Displacement(dt float64) pixel.Vec {
	p.velocityY -= p.Gravity * dt
	return pixel.V(p.velocityX*dt, p.velocityY*dt)
}

// IsStopped returns true is the item is not moving
func (p *Physics) IsStopped() bool {
	return p.velocityX == 0
}

func (p *Physics) IsMovingUp() bool {
	return p.velocityY > 0
}

func (p *Physics) StopMovingY() {
	p.velocityY = 0
}

func (p *Physics) StopMovingX() {
	p.velocityX = 0
}

// IsMovingDown returns true if the item has a negative velocity
func (p *Physics) IsMovingDown() bool {
	return p.velocityY < 0
}

func (p *Physics) VelocityY() float64 {
	return p.velocityY
}
