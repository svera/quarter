package level

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/svera/quarter/physic"
)

// Level holds the information needed to build a level
type Level struct {
	Layers []Layer
	imd    *imdraw.IMDraw
	debug  bool
	Extra  interface{}
}

// Layer contains the different structs a layer can hold and show on screen
type Layer struct {
	image  *pixel.Sprite
	Grid   *Grid
	Bounds []physic.Shaper
	Extra  interface{}
}

// Draw renders the level following layers order
func (l *Level) Draw(target pixel.Target) {
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
		if layer.Bounds != nil && l.debug {
			for _, b := range layer.Bounds {
				b.Shape().Draw(color.RGBA{0, 0, 255, 16}, l.imd, target)
			}
		}
	}
}

func (l *Level) SetDebug(imd *imdraw.IMDraw) {
	l.imd = imd
	l.debug = true
}
