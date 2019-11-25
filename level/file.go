package level

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/faiface/pixel"
	"github.com/svera/quarter"
	"github.com/svera/quarter/physic"
)

// LevelsFile is the serialized form of Levels
type LevelsFile struct {
	Version string
	Levels  []struct {
		Name   string
		Layers []struct {
			Name  string
			Image struct {
				Path string
			}
			Grid struct {
				Width  int
				Height int
				Assets GridAssets
				Tiles  []struct {
					Asset int
					X     float64
					Y     float64
					Extra interface{}
				}
				Extra interface{}
			}
			Bounds []struct {
				Type       string
				Dimensions json.RawMessage
			}
			Extra interface{}
		}
		Extra interface{}
	}
}

type GridAssets struct {
	Path     string
	Quantity int
	OffsetX  float64 `json:"offset_x"`
	OffsetY  float64 `json:"offset_y"`
	Width    float64
	Height   float64
}

type BoundRectDimensions struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func Load(r io.Reader) ([]Level, error) {
	data := LevelsFile{}
	levels := make([]Level, len(data.Levels))

	err := json.NewDecoder(r).Decode(&data)
	if err != nil {
		return nil, err
	}
	if data.Version != "1" {
		return nil, fmt.Errorf("Version not supported")
	}

	for _, currentLevel := range data.Levels {
		level := Level{}
		for i, currentLayer := range currentLevel.Layers {
			level.Layers = append(level.Layers, Layer{})
			if path := strings.TrimSpace(currentLayer.Image.Path); path != "" {
				img, err := quarter.LoadPicture(path)
				if err != nil {
					return nil, err
				}
				level.Layers[i].Image = pixel.NewSprite(img, img.Bounds())
			}
			if currentLayer.Grid.Width != 0 || currentLayer.Grid.Height != 0 {
				level.Layers[i].Grid = &Grid{
					Width:       currentLayer.Grid.Width,
					Height:      currentLayer.Grid.Height,
					assetWidth:  currentLayer.Grid.Assets.Width,
					assetHeight: currentLayer.Grid.Assets.Height,
				}
				level.Layers[i].Grid.Assets, err = loadGridAssets(currentLayer.Grid.Assets)
				if err != nil {
					return nil, err
				}
				level.Layers[i].Grid.Tiles = make([]GridTile, len(currentLayer.Grid.Tiles))
				for k := range currentLayer.Grid.Tiles {
					tl := GridTile{
						Asset:  currentLayer.Grid.Tiles[k].Asset,
						Coords: pixel.V(currentLayer.Grid.Tiles[k].X, currentLayer.Grid.Tiles[k].Y),
					}
					level.Layers[i].Grid.Tiles[k] = tl
				}
			}
			if len(currentLayer.Bounds) > 0 {
				level.Layers[i].Bounds = make([]physic.Shaper, len(currentLayer.Bounds))
				for j, bound := range currentLayer.Bounds {
					if bound.Type == "box" {
						dimensions := BoundRectDimensions{}
						err := json.Unmarshal(bound.Dimensions, &dimensions)
						if err != nil {
							return nil, fmt.Errorf("Bounds dimensions wrongly formatted")
						}
						level.Layers[i].Bounds[j] = physic.NewBoundingBox(dimensions.X, dimensions.Y, dimensions.Width, dimensions.Height)
					}
				}
			}
		}
		levels = append(levels, level)
	}

	return levels, nil
}

func loadGridAssets(assets GridAssets) ([]*pixel.Sprite, error) {
	img, err := quarter.LoadPicture(assets.Path)
	if err != nil {
		return nil, err
	}

	sprites := make([]*pixel.Sprite, assets.Quantity)
	for j := 0; j < assets.Quantity; j++ {
		x := assets.Width * float64(j)
		sprite := pixel.NewSprite(
			img,
			pixel.R(
				x,
				assets.OffsetY,
				x+assets.Width,
				assets.OffsetY+assets.Height,
			),
		)
		sprites[j] = sprite
	}
	return sprites, nil
}
