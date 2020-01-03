package animation_test

import (
	"bytes"
	"testing"

	"github.com/faiface/pixel"
	"github.com/svera/quarter/animation"
)

func TestLoad(t *testing.T) {
	t.Run("Only valid JSON is supported", func(t *testing.T) {
		levelData := []byte(``)
		r := bytes.NewReader(levelData)
		if _, err := animation.LoadBoundedAnimation(r, pixel.V(0, 0)); err == nil {
			t.Errorf("An invalid JSON file must return error")
		}
	})

	t.Run("Only version 1 is supported", func(t *testing.T) {
		levelData := []byte(`{"version": "1", "anims": [{}]}`)
		r := bytes.NewReader(levelData)
		if _, err := animation.LoadBoundedAnimation(r, pixel.V(0, 0)); err != nil {
			t.Errorf("Valid animation data is not loaded")
		}
		levelData = []byte(`{"version": "2", "anims": [{}]}`)
		r = bytes.NewReader(levelData)
		if _, err := animation.LoadBoundedAnimation(r, pixel.V(0, 0)); err == nil {
			t.Errorf("Invalid animation data is loaded")
		}
	})

}
