package level

import (
	"github.com/faiface/pixel"
)

// Grid divides the screen in tiles of tileWidth * tileHeight size
type Grid struct {
	TileWidth  float64
	TileHeight float64
	Assets     []*pixel.Sprite
	Tiles      []GridTile
}

// GridTile holds data of a single tile
type GridTile struct {
	Asset  int
	Coords pixel.Vec
}

// NewGrid returns a new Grid instance
func NewGrid(w, h float64) *Grid {
	return &Grid{
		TileWidth:  w,
		TileHeight: h,
	}
}

// ToPixels transforms grid coordinates of a tile to its center coordinates in pixels.
func (g *Grid) ToPixels(coords pixel.Vec) pixel.Vec {
	return pixel.V(
		(coords.X*g.TileWidth)+g.TileWidth/2,
		(coords.Y*g.TileHeight)+g.TileHeight/2,
	)
}
