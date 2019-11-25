package level

import "github.com/faiface/pixel"

type Grid struct {
	Width       int
	Height      int
	assetWidth  float64
	assetHeight float64
	Assets      []*pixel.Sprite
	Tiles       []GridTile
	Extra       interface{}
}

type GridTile struct {
	Asset  int
	Coords pixel.Vec
	Extra  interface{}
}

func NewGrid() *Grid {
	return &Grid{}
}

func (g *Grid) ToPixels(coords pixel.Vec) pixel.Vec {
	//return pixel.V(coords.X*g.Assets[0].Frame().W(), coords.Y*g.Assets[0].Frame().H())
	return pixel.V(
		(coords.X*g.assetWidth)+g.assetWidth/2,
		(coords.Y*g.assetHeight)+g.assetHeight/2,
	)
}
