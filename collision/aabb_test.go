package collision_test

// import (
// 	"testing"

// 	"github.com/svera/animated-sprite/pkg/quarter"
// )

// func TestCollide(t *testing.T) {
// 	rect1 := quarter.BoundingBox{X: 5, Y: 5, Width: 50, Height: 50}
// 	rect2 := quarter.BoundingBox{X: 20, Y: 10, Width: 10, Height: 10}

// 	t.Run("Collision is detected", func(t *testing.T) {
// 		if !quarter.Collide(rect1, rect2) {
// 			t.Errorf("Rect1 collides with Rect2 but no collision is detected")
// 		}
// 	})

// 	t.Run("No collision is detected", func(t *testing.T) {
// 		rect1.Width = 10
// 		if quarter.Collide(rect1, rect2) {
// 			t.Errorf("Rect1 does not collide with Rect2 but a collision is detected")
// 		}
// 	})
// }

// func TestResolve(t *testing.T) {
// 	/*
// 		+-----+
// 		| r1  |
// 		|     |
// 		+-----+
// 				\
// 				 v
// 				 	+-----+
// 					| r2  |
// 					|     |
// 					+-----+
// 	*/
// 	t.Run("rect1 moves right and down and collides with rect2", func(t *testing.T) {
// 		rect1 := quarter.BoundingBox{X: 5, Y: 20, Width: 10, Height: 10}
// 		rect2 := quarter.BoundingBox{X: 20, Y: 5, Width: 10, Height: 10}

// 		expectedMaxDx := 4.0
// 		expectedMaxDy := -4.0
// 		collide, maxdx, maxdy := quarter.Resolve(rect1, rect2, 10, -10)
// 		if !collide || maxdx != expectedMaxDx || maxdy != expectedMaxDy {
// 			t.Errorf("Wrong resolution values, expected %v, %f, %f, got %v, %f, %f", true, expectedMaxDx, expectedMaxDy, collide, maxdx, maxdy)
// 		}
// 	})

// 	/*
// 		+-----+
// 		| r2  |
// 		|     |
// 		+-----+
// 				^
// 		          \
// 				 	+-----+
// 					| r1  |
// 					|     |
// 					+-----+
// 	*/
// 	t.Run("rect1 moves left and up and collides with rect2", func(t *testing.T) {
// 		rect1 := quarter.BoundingBox{X: 20, Y: 5, Width: 10, Height: 10}
// 		rect2 := quarter.BoundingBox{X: 5, Y: 20, Width: 10, Height: 10}

// 		expectedMaxDx := -4.0
// 		expectedMaxDy := 4.0
// 		collide, maxdx, maxdy := quarter.Resolve(rect1, rect2, -10, 10)
// 		if !collide || maxdx != expectedMaxDx || maxdy != expectedMaxDy {
// 			t.Errorf("Wrong resolution values, expected %v, %f, %f, got %v, %f, %f", true, expectedMaxDx, expectedMaxDy, collide, maxdx, maxdy)
// 		}
// 	})

// 	/*
// 		+-----+    		+-----+
// 		| r1  |         | r2  |
// 		|     |  --->   |     |
// 		+-----+         +-----+
// 	*/
// 	t.Run("rect1 moves right and collides with rect2", func(t *testing.T) {
// 		rect1 := quarter.BoundingBox{X: 5, Y: 5, Width: 10, Height: 10}
// 		rect2 := quarter.BoundingBox{X: 20, Y: 5, Width: 10, Height: 10}

// 		expectedMaxDx := 4.0
// 		expectedMaxDy := 0.0
// 		collide, maxdx, maxdy := quarter.Resolve(rect1, rect2, 10, 0)
// 		if !collide || maxdx != expectedMaxDx || maxdy != expectedMaxDy {
// 			t.Errorf("Wrong resolution values, expected %v, %f, %f, got %v, %f, %f", true, expectedMaxDx, expectedMaxDy, collide, maxdx, maxdy)
// 		}
// 	})

// 	/*
// 		 +-----+
// 		 | r1  |
// 		 |     |
// 		 +-----+
// 		    |
// 		    v
// 		+--------+
// 		|   r2   |
// 		|        |
// 		+--------+
// 	*/
// 	t.Run("rect1 moves down and collides with rect2", func(t *testing.T) {
// 		rect1 := quarter.BoundingBox{X: 20, Y: 20, Width: 10, Height: 10}
// 		rect2 := quarter.BoundingBox{X: 10, Y: 5, Width: 30, Height: 10}

// 		expectedMaxDx := 0.0
// 		expectedMaxDy := -4.0
// 		collide, maxdx, maxdy := quarter.Resolve(rect1, rect2, 0, -10)
// 		if !collide || maxdx != expectedMaxDx || maxdy != expectedMaxDy {
// 			t.Errorf("Wrong resolution values, expected %v, %f, %f, got %v, %f, %f", true, expectedMaxDx, expectedMaxDy, collide, maxdx, maxdy)
// 		}
// 	})

// 	/*
// 		 +-------+
// 		 |  r2   |
// 		 |       |
// 		 +-------+
// 			^
// 		    |
// 		  +-----+
// 		  | r1  |
// 		  |     |
// 		  +-----+
// 	*/
// 	t.Run("rect1 moves up and collides with rect2", func(t *testing.T) {
// 		rect1 := quarter.BoundingBox{X: 20, Y: 5, Width: 10, Height: 10}
// 		rect2 := quarter.BoundingBox{X: 10, Y: 20, Width: 30, Height: 10}

// 		expectedMaxDx := 0.0
// 		expectedMaxDy := 4.0
// 		collide, maxdx, maxdy := quarter.Resolve(rect1, rect2, 0, 10)
// 		if !collide || maxdx != expectedMaxDx || maxdy != expectedMaxDy {
// 			t.Errorf("Wrong resolution values, expected %v, %f, %f, got %v, %f, %f", true, expectedMaxDx, expectedMaxDy, collide, maxdx, maxdy)
// 		}
// 	})

// }
