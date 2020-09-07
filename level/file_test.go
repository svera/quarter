package level_test

import (
	"bytes"
	"testing"

	"github.com/svera/quarter/level"
)

func TestLoad(t *testing.T) {
	t.Run("Only valid JSON is supported", func(t *testing.T) {
		levelData := []byte(``)
		r := bytes.NewReader(levelData)
		if _, err := level.Deserialize(r); err == nil {
			t.Errorf("An invalid JSON file must return error")
		}
	})

	t.Run("Only version 1 is supported", func(t *testing.T) {
		levelData := []byte(`{"version": "1", "levels": {"name": {}}}`)
		r := bytes.NewReader(levelData)
		if _, err := level.Deserialize(r); err != nil {
			t.Errorf("Valid levels data is not loaded")
		}
		levelData = []byte(`{"version": "2", "levels": {"name": {}}}`)
		r = bytes.NewReader(levelData)
		if _, err := level.Deserialize(r); err == nil {
			t.Errorf("Invalid levels data is loaded")
		}
	})

	t.Run("Levels file must have at least one level", func(t *testing.T) {
		levelData := []byte(`{"version": "1"}`)
		r := bytes.NewReader(levelData)
		if _, err := level.Deserialize(r); err == nil {
			t.Errorf("Levels file must have at least one level")
		}
	})
}
