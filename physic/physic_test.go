package physic_test

import (
	"reflect"
	"testing"

	"github.com/faiface/pixel"
	"github.com/svera/quarter/physic"
)

func TestAccelerateX(t *testing.T) {
	var testValues = []struct {
		testName          string
		dir               float64
		dt                float64
		expectedVelocityX float64
	}{
		{"Accelerate to the right", physic.DirectionRight, 0.5, 5},
		{"Accelerate to the left", physic.DirectionLeft, 0.5, -5},
	}
	for _, tt := range testValues {
		t.Run(tt.testName, func(t *testing.T) {
			phys := physic.NewPhysics(physic.Params{
				MaxVelocityX:  10,
				AccelerationX: 10,
			})
			phys.AccelerateX(tt.dir, tt.dt)
			if phys.VelocityX() != tt.expectedVelocityX {
				t.Errorf("got %f, want %f", phys.VelocityX(), tt.expectedVelocityX)
			}
		})
	}
}

func TestDecelerateX(t *testing.T) {
	var testValues = []struct {
		testName          string
		dt                float64
		expectedVelocityX float64
	}{
		{"Decelarate ", 0.5, 0},
		{"Velocity after decelaration cannot be less than 0", 0.5, 0},
	}
	phys := physic.NewPhysics(physic.Params{
		MaxVelocityX:  10,
		AccelerationX: 10,
	})
	phys.AccelerateX(physic.DirectionRight, 0.5)
	for _, tt := range testValues {
		t.Run(tt.testName, func(t *testing.T) {
			phys.DecelerateX(tt.dt)
			if phys.VelocityX() != tt.expectedVelocityX {
				t.Errorf("got %f, want %f", phys.VelocityX(), tt.expectedVelocityX)
			}
		})
	}
}

func TestDisplacement(t *testing.T) {
	var testValues = []struct {
		testName             string
		dt                   float64
		gravity              float64
		expectedDisplacement pixel.Vec
	}{
		{"Displacement without gravity ", 0.5, 0, pixel.V(2.5, 1.25)},
		{"Displacement with gravity", 0.5, 1, pixel.V(2.5, 1)},
	}

	phys := physic.NewPhysics(physic.Params{
		MaxVelocityX:  10,
		MaxVelocityY:  10,
		AccelerationX: 10,
		AccelerationY: 5,
	})
	phys.AccelerateX(physic.DirectionRight, 0.5)
	phys.AccelerateY(physic.DirectionUp, 0.5)
	for _, tt := range testValues {
		t.Run(tt.testName, func(t *testing.T) {
			phys.Gravity = tt.gravity
			dis := phys.Displacement(0.5)
			if !reflect.DeepEqual(dis, tt.expectedDisplacement) {
				t.Errorf("Expected displacement of %v, got %v", tt.expectedDisplacement, dis)
			}
		})
	}
}

func TestSetVelocityX(t *testing.T) {
	var testValues = []struct {
		testName          string
		velocityX         float64
		expectedVelocityX float64
	}{
		{"Velocity X lower than max", 5, 5},
		{"Velocity X greater than max", 15, 10},
	}

	phys := physic.NewPhysics(physic.Params{
		MaxVelocityX: 10,
	})
	for _, tt := range testValues {
		t.Run(tt.testName, func(t *testing.T) {
			phys.SetVelocityX(tt.velocityX)
			if phys.VelocityX() != tt.expectedVelocityX {
				t.Errorf("Expected velocity X is %f, got %f", tt.expectedVelocityX, phys.VelocityX())
			}
		})
	}
}

func TestSetVelocityY(t *testing.T) {
	var testValues = []struct {
		testName          string
		velocityY         float64
		expectedVelocityY float64
	}{
		{"Velocity Y lower than max", 5, 5},
		{"Velocity Y greater than max", 15, 10},
	}

	phys := physic.NewPhysics(physic.Params{
		MaxVelocityY: 10,
	})
	for _, tt := range testValues {
		t.Run(tt.testName, func(t *testing.T) {
			phys.SetVelocityY(tt.velocityY)
			if phys.VelocityY() != tt.expectedVelocityY {
				t.Errorf("Expected velocity Y is %f, got %f", tt.expectedVelocityY, phys.VelocityY())
			}
		})
	}
}
