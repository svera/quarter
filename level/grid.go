package level

import "github.com/faiface/pixel"

type Grid struct {
	Width  int
	Height int
	sprite *pixel.Sprite
	batch  *pixel.Batch
	Sheet  pixel.Picture
	Assets []pixel.Rect
	Tiles  []GridTile
	Extra  interface{}
}

type GridTile struct {
	Asset  int
	Coords pixel.Vec
	Extra  interface{}
}

func NewGrid(width, height int, sheet pixel.Picture) *Grid {
	return &Grid{
		Width:  width,
		Height: height,
		Sheet:  sheet,
		sprite: pixel.NewSprite(sheet, pixel.Rect{}),
		batch:  pixel.NewBatch(&pixel.TrianglesData{}, sheet),
	}
}

func (g *Grid) ToPixels(coords pixel.Vec) pixel.Vec {
	return pixel.V(coords.X*g.Assets[0].W(), coords.Y*g.Assets[0].H())
}
