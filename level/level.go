package level

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/svera/quarter/physic"
)

// Level holds the information needed
type Level struct {
	Layers []Layer
	imd    *imdraw.IMDraw
	debug  bool
	Extra  interface{}
}

type Layer struct {
	Image  *pixel.Picture
	Grid   *Grid
	Bounds []physic.Shaper
	Extra  interface{}
}

func (l *Level) Draw(target pixel.Target) {
	for _, layer := range l.Layers {
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
