package collision_test

import (
	"reflect"
	"testing"

	"github.com/faiface/pixel"
	"github.com/svera/quarter/collision"
)

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
		rect1 := &collision.BoundingBox{pixel.Rect{Min: pixel.V(5, 20), Max: pixel.V(15, 30)}}
		rect2 := &collision.BoundingBox{pixel.Rect{Min: pixel.V(20, 5), Max: pixel.V(30, 15)}}

		expectedSolution := collision.Solution{
			CollisionAxis: collision.AxisX,
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
		rect1 := &collision.BoundingBox{pixel.Rect{Min: pixel.V(20, 5), Max: pixel.V(30, 15)}}
		rect2 := &collision.BoundingBox{pixel.Rect{Min: pixel.V(5, 20), Max: pixel.V(15, 30)}}

		expectedSolution := collision.Solution{
			CollisionAxis: collision.AxisX,
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
		rect1 := &collision.BoundingBox{pixel.Rect{Min: pixel.V(5, 5), Max: pixel.V(15, 15)}}
		rect2 := &collision.BoundingBox{pixel.Rect{Min: pixel.V(20, 5), Max: pixel.V(30, 15)}}

		expectedSolution := collision.Solution{
			CollisionAxis: collision.AxisX,
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
		rect1 := &collision.BoundingBox{pixel.Rect{Min: pixel.V(20, 20), Max: pixel.V(30, 30)}}
		rect2 := &collision.BoundingBox{pixel.Rect{Min: pixel.V(10, 5), Max: pixel.V(40, 15)}}

		expectedSolution := collision.Solution{
			CollisionAxis: collision.AxisY,
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
		rect1 := &collision.BoundingBox{pixel.Rect{Min: pixel.V(20, 5), Max: pixel.V(30, 15)}}
		rect2 := &collision.BoundingBox{pixel.Rect{Min: pixel.V(10, 20), Max: pixel.V(40, 30)}}

		expectedSolution := collision.Solution{
			CollisionAxis: collision.AxisY,
			Object:        rect2,
			Distance:      pixel.V(0, 5),
		}
		sol := rect1.Resolve(pixel.V(0, 10), rect2)
		if !reflect.DeepEqual(sol, expectedSolution) {
			t.Errorf("Wrong resolution values, expected %v, got %v", expectedSolution, sol)
		}
	})
}
