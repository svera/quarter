package physic_test

import (
	"reflect"
	"testing"

	"github.com/faiface/pixel"
	"github.com/svera/quarter/physic"
)

func TestAccelerate(t *testing.T) {
	var testValues = []struct {
		testName         string
		axis             int
		dir              float64
		dt               float64
		expectedVelocity float64
	}{
		{"Accelerate to the right", physic.AxisX, physic.DirectionRight, 0.5, 5},
		{"Accelerate to the left", physic.AxisX, physic.DirectionLeft, 0.5, -5},
	}
	for _, tt := range testValues {
		t.Run(tt.testName, func(t *testing.T) {
			phys := physic.NewPhysics(
				physic.Params{
					MaxVelocity:  [2]float64{10, 0},
					Acceleration: [2]float64{10, 0},
				},
			)
			phys.Accelerate(tt.axis, tt.dir, tt.dt)
			if phys.Velocity(tt.axis) != tt.expectedVelocity {
				t.Errorf("got %f, want %f", phys.Velocity(tt.axis), tt.expectedVelocity)
			}
		})
	}
}

func TestDecelerate(t *testing.T) {
	var testValues = []struct {
		testName         string
		axis             int
		dt               float64
		expectedVelocity float64
	}{
		{"Decelerate ", physic.AxisX, .5, 0},
		{"Velocity after deceleration cannot be less than 0", physic.AxisX, .5, 0},
	}
	for _, tt := range testValues {
		t.Run(tt.testName, func(t *testing.T) {
			phys := physic.NewPhysics(physic.Params{
				MaxVelocity:  [2]float64{10, 0},
				Acceleration: [2]float64{10, 0},
			})
			phys.Accelerate(tt.axis, physic.DirectionRight, .5)
			phys.Decelerate(tt.axis, tt.dt)
			if phys.Velocity(tt.axis) != tt.expectedVelocity {
				t.Errorf("got %f, want %f", phys.Velocity(tt.axis), tt.expectedVelocity)
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
		MaxVelocity:  [2]float64{10, 10},
		Acceleration: [2]float64{10, 5},
	})
	phys.Accelerate(physic.AxisX, physic.DirectionRight, 0.5)
	phys.Accelerate(physic.AxisY, physic.DirectionUp, 0.5)
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

func TestSetVelocity(t *testing.T) {
	var testValues = []struct {
		testName         string
		axis             int
		velocity         float64
		expectedVelocity float64
	}{
		{"Velocity X lower than max", physic.AxisX, 5, 5},
		{"Velocity X greater than max", physic.AxisX, 15, 10},
		{"Velocity Y lower than max", physic.AxisX, 5, 5},
		{"Velocity Y greater than max", physic.AxisX, 15, 10},
	}

	phys := physic.NewPhysics(physic.Params{
		MaxVelocity: [2]float64{10, 10},
	})
	for _, tt := range testValues {
		t.Run(tt.testName, func(t *testing.T) {
			phys.SetVelocity(tt.axis, tt.velocity)
			if phys.Velocity(tt.axis) != tt.expectedVelocity {
				t.Errorf("Expected velocity is %f, got %f", tt.expectedVelocity, phys.Velocity(tt.axis))
			}
		})
	}
}
