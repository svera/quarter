package level_test

import (
	"reflect"
	"testing"

	"github.com/faiface/pixel"
	"github.com/svera/quarter/level"
)

func TestToPixels(t *testing.T) {
	g := level.NewGrid(32, 32)
	gridCoords := pixel.V(0, 0)
	screenCoords := g.ToPixels(gridCoords)
	expectedScreenCoords := pixel.V(16, 16)
	if !reflect.DeepEqual(screenCoords, expectedScreenCoords) {
		t.Errorf("Grid coords %v must translate to screen coords %v, got %v", gridCoords, expectedScreenCoords, screenCoords)
	}
}
