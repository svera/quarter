package level

import (
	"github.com/faiface/pixel"
)

// Level holds the information needed to build a level
type Level struct {
	Limits pixel.Rect
	Layers map[string]Layer
}

// Layer contains the different structs a layer can hold and show on screen
type Layer struct {
	image *pixel.Sprite
	Grid  *Grid
}

// Draw renders the level following layers order
func (l Level) Draw(target pixel.Target) {
	for _, layer := range l.Layers {
		if layer.image != nil {
			layer.image.Draw(target, pixel.IM.Moved(layer.image.Frame().Min))
		}
		if layer.Grid != nil {
			for _, t := range layer.Grid.Tiles {
				pixelCoords := layer.Grid.ToPixels(t.Coords)
				layer.Grid.Assets[t.Asset].Draw(target, pixel.IM.Moved(pixelCoords))
			}
		}
	}
}
