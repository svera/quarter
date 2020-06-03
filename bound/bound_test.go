package bound_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/faiface/pixel"
	"github.com/svera/quarter/bound"
)

func TestCollidesWithBoundingBox(t *testing.T) {
	rect1 := &bound.Box{pixel.Rect{Min: pixel.V(5, 5), Max: pixel.V(55, 55)}}
	rect2 := &bound.Box{pixel.Rect{Min: pixel.V(20, 10), Max: pixel.V(30, 20)}}

	t.Run("Collision is detected", func(t *testing.T) {
		if !rect1.Collides(rect2) {
			t.Errorf("Rect1 collides with Rect2 but no collision is detected")
		}
	})

	t.Run("No collision is detected", func(t *testing.T) {
		rect1.Max.X = 15
		if rect1.Collides(rect2) {
			t.Errorf("Rect1 does not collide with Rect2 but a collision is detected")
		}
	})
}

func TestResolveBoxAgainstBox(t *testing.T) {
	/*
		+-----+
		| r1  |
		|     |
		+-----+
				\
				 v
				 	+-----+
					| r2  |
					|     |
					+-----+
	*/
	t.Run("rect1 moves right and down and collides with rect2", func(t *testing.T) {
		rect1 := &bound.Box{pixel.Rect{Min: pixel.V(5, 20), Max: pixel.V(15, 30)}}
		rect2 := &bound.Box{pixel.Rect{Min: pixel.V(20, 5), Max: pixel.V(30, 15)}}

		expectedSolution := bound.Solution{
			CollisionAxis: bound.AxisX,
			Object:        rect2,
			Distance:      pixel.V(5, 0),
		}
		sol := rect1.Resolve(pixel.V(10, -10), rect2)
		if !reflect.DeepEqual(sol, expectedSolution) {
			t.Errorf("Wrong resolution values, expected %v, got %v", expectedSolution, sol)
		}
	})

	/*
		+-----+
		| r2  |
		|     |
		+-----+
				^
		          \
				 	+-----+
					| r1  |
					|     |
					+-----+
	*/
	t.Run("rect1 moves left and up and collides with rect2", func(t *testing.T) {
		rect1 := &bound.Box{pixel.Rect{Min: pixel.V(20, 5), Max: pixel.V(30, 15)}}
		rect2 := &bound.Box{pixel.Rect{Min: pixel.V(5, 20), Max: pixel.V(15, 30)}}

		expectedSolution := bound.Solution{
			CollisionAxis: bound.AxisX,
			Object:        rect2,
			Distance:      pixel.V(-5, 0),
		}
		sol := rect1.Resolve(pixel.V(-10, 10), rect2)
		if !reflect.DeepEqual(sol, expectedSolution) {
			t.Errorf("Wrong resolution values, expected %v, got %v", expectedSolution, sol)
		}
	})

	/*
		+-----+    		+-----+
		| r1  |         | r2  |
		|     |  --->   |     |
		+-----+         +-----+
	*/
	t.Run("rect1 moves right and collides with rect2", func(t *testing.T) {
		rect1 := &bound.Box{pixel.Rect{Min: pixel.V(5, 5), Max: pixel.V(15, 15)}}
		rect2 := &bound.Box{pixel.Rect{Min: pixel.V(20, 5), Max: pixel.V(30, 15)}}

		expectedSolution := bound.Solution{
			CollisionAxis: bound.AxisX,
			Object:        rect2,
			Distance:      pixel.V(5, 0),
		}
		sol := rect1.Resolve(pixel.V(10, 0), rect2)
		if !reflect.DeepEqual(sol, expectedSolution) {
			t.Errorf("Wrong resolution values, expected %v, got %v", expectedSolution, sol)
		}
	})

	/*
		 +-----+
		 | r1  |
		 |     |
		 +-----+
		    |
		    v
		+--------+
		|   r2   |
		|        |
		+--------+
	*/
	t.Run("rect1 moves down and collides with rect2", func(t *testing.T) {
		rect1 := &bound.Box{pixel.Rect{Min: pixel.V(20, 20), Max: pixel.V(30, 30)}}
		rect2 := &bound.Box{pixel.Rect{Min: pixel.V(10, 5), Max: pixel.V(40, 15)}}

		expectedSolution := bound.Solution{
			CollisionAxis: bound.AxisY,
			Object:        rect2,
			Distance:      pixel.V(0, -5),
		}
		sol := rect1.Resolve(pixel.V(0, -10), rect2)
		if !reflect.DeepEqual(sol, expectedSolution) {
			t.Errorf("Wrong resolution values, expected %v, got %v", expectedSolution, sol)
		}
	})

	/*
		 +-------+
		 |  r2   |
		 |       |
		 +-------+
			^
		    |
		  +-----+
		  | r1  |
		  |     |
		  +-----+
	*/
	t.Run("rect1 moves up and collides with rect2", func(t *testing.T) {
		rect1 := &bound.Box{pixel.Rect{Min: pixel.V(20, 5), Max: pixel.V(30, 15)}}
		rect2 := &bound.Box{pixel.Rect{Min: pixel.V(10, 20), Max: pixel.V(40, 30)}}

		expectedSolution := bound.Solution{
			CollisionAxis: bound.AxisY,
			Object:        rect2,
			Distance:      pixel.V(0, 5),
		}
		sol := rect1.Resolve(pixel.V(0, 10), rect2)
		if !reflect.DeepEqual(sol, expectedSolution) {
			t.Errorf("Wrong resolution values, expected %v, got %v", expectedSolution, sol)
		}
	})
}

func TestLoad(t *testing.T) {
	t.Run("Only valid JSON is supported", func(t *testing.T) {
		levelData := []byte(``)
		r := bytes.NewReader(levelData)
		if _, err := bound.Deserialize(r); err == nil {
			t.Errorf("An invalid JSON file must return error")
		}
	})

	t.Run("Only version 1 is supported", func(t *testing.T) {
		levelData := []byte(`{"version": "1", "bounds": [{}]}`)
		r := bytes.NewReader(levelData)
		if _, err := bound.Deserialize(r); err != nil {
			t.Errorf("Valid collision data is not loaded")
		}
		levelData = []byte(`{"version": "2", "bounds": [{}]}`)
		r = bytes.NewReader(levelData)
		if _, err := bound.Deserialize(r); err == nil {
			t.Errorf("Invalid collision data is loaded")
		}
	})

}
