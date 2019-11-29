package level

import (
	"github.com/faiface/pixel"
	"github.com/svera/quarter/collision"
)

// Grid divides the screen in tiles of tileWidth * tileHeight size
type Grid struct {
	tileWidth  float64
	tileHeight float64
	Assets     []*pixel.Sprite
	Tiles      []GridTile
	Extra      interface{}
}

// GridTile holds data of a single tile
type GridTile struct {
	Asset  int
	Coords pixel.Vec
	Extra  interface{}
}

// NewGrid returns a new Grid instance
func NewGrid() *Grid {
	return &Grid{}
}

// ToPixels transforms grid coordinates of a tile to its center coordinates in pixels.
func (g *Grid) ToPixels(coords pixel.Vec) pixel.Vec {
	return pixel.V(
		(coords.X*g.tileWidth)+g.tileWidth/2,
		(coords.Y*g.tileHeight)+g.tileHeight/2,
	)
}

func (g *Grid) TileBoundingBox(coords pixel.Vec) *collision.BoundingBox {
	min := pixel.V(coords.X*g.tileWidth, coords.Y*g.tileHeight)
	max := pixel.V(min.X+g.tileWidth, min.Y+g.tileHeight)
	return collision.NewBoundingBox(min, max)
}
