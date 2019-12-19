package level_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/svera/quarter/level"
)

func TestLoad(t *testing.T) {
	t.Run("Only valid JSON is supported", func(t *testing.T) {
		levelData := []byte(``)
		r := bytes.NewReader(levelData)
		if _, err := level.Load(r); err == nil {
			t.Errorf("An invalid JSON file must return error")
		}
	})

	t.Run("Only version 1 is supported", func(t *testing.T) {
		levelData := []byte(`{"version": "1", "levels": [{}]}`)
		r := bytes.NewReader(levelData)
		if _, err := level.Load(r); err != nil {
			t.Errorf("Valid levels data is not loaded")
		}
		levelData = []byte(`{"version": "2", "levels": [{}]}`)
		r = bytes.NewReader(levelData)
		if _, err := level.Load(r); err == nil {
			t.Errorf("Invalid levels data is loaded")
		}
	})

	t.Run("Levels file must have at least one level", func(t *testing.T) {
		levelData := []byte(`{"version": "1"}`)
		r := bytes.NewReader(levelData)
		if _, err := level.Load(r); err == nil {
			t.Errorf("Levels file must have at least one level")
		}
	})

	t.Run("Levels file with unsupported bound type must return error", func(t *testing.T) {
		levelData := []byte(`
			{
				"version": "1",
				"levels": [
					{
						"layers": [
							{
								"bounds": [
									{
										"type": "hexagon"
									}
								]
							}
						]
					}
				]
			}
		`)
		r := bytes.NewReader(levelData)
		if _, err := level.Load(r); err.Error() != fmt.Sprintf(level.ErrorBoundTypeNotSupported, "hexagon") {
			t.Errorf("Levels file with unsupported bound type must return error")
		}
	})

	t.Run("Levels file with supported bound type but with wrong values must return error", func(t *testing.T) {
		levelData := []byte(`
			{
				"version": "1",
				"levels": [
					{
						"layers": [
							{
								"bounds": [
									{
										"type": "box"
									}
								]
							}
						]
					}
				]
			}
		`)
		r := bytes.NewReader(levelData)
		if _, err := level.Load(r); err.Error() != fmt.Sprintf(level.ErrorWrongBoundValues) {
			t.Errorf("Levels file with wrong bound values must return error")
		}
	})

}
