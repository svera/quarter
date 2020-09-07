package level

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/faiface/pixel"
	"github.com/svera/quarter"
)

// Returned errors
const (
	ErrorWrongFileFormat     = "Loaded levels file is not a valid JSON"
	ErrorVersionNotSupported = "Version \"%s\" not supported"
	ErrorNoLevels            = "Levels file must have at least one level declared, none found"
)

// LevelsFile is the serialized form of Levels
type LevelsFile struct {
	Version string
	Levels  map[string]struct {
		Limits struct {
			Min pixel.Vec
			Max pixel.Vec
		}
		Layers map[string]struct {
			Image struct {
				Path string
			}
			Grid struct {
				Assets GridAssets
				Tiles  []struct {
					Asset int
					X     float64
					Y     float64
				}
			}
		}
	}
}

type GridAssets struct {
	Path     string
	Quantity int
	Offset   struct {
		X float64
		Y float64
	}
	Width  float64
	Height float64
}

// Deserialize validates a levels file and returns its information as a []Level
func Deserialize(r io.Reader) (map[string]Level, error) {
	data := LevelsFile{}
	err := json.NewDecoder(r).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Version != "1" {
		return nil, fmt.Errorf(ErrorVersionNotSupported, data.Version)
	}

	if len(data.Levels) == 0 {
		return nil, fmt.Errorf(ErrorNoLevels)
	}

	levels := make(map[string]Level, len(data.Levels))

	for levelName, levelData := range data.Levels {
		level := Level{
			Limits: pixel.R(levelData.Limits.Min.X, levelData.Limits.Min.Y, levelData.Limits.Max.X, levelData.Limits.Max.Y),
			Layers: make(map[string]Layer),
		}
		for layerName, layerData := range levelData.Layers {
			layer := Layer{}
			if path := strings.TrimSpace(layerData.Image.Path); path != "" {
				img, err := quarter.LoadPicture(path)
				if err != nil {
					return nil, err
				}
				layer.image = pixel.NewSprite(img, img.Bounds())
			}
			if len(layerData.Grid.Tiles) != 0 {
				layer.Grid = &Grid{
					TileWidth:  layerData.Grid.Assets.Width,
					TileHeight: layerData.Grid.Assets.Height,
				}
				layer.Grid.Assets, err = loadGridAssets(layerData.Grid.Assets)
				if err != nil {
					return nil, err
				}
				layer.Grid.Tiles = make([]GridTile, len(layerData.Grid.Tiles))
				for k, val := range layerData.Grid.Tiles {
					tl := GridTile{
						Asset:  val.Asset,
						Coords: pixel.V(val.X, val.Y),
					}
					layer.Grid.Tiles[k] = tl
				}
			}
			level.Layers[layerName] = layer
		}
		levels[levelName] = level
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
				assets.Offset.Y,
				x+assets.Width,
				assets.Offset.Y+assets.Height,
			),
		)
		sprites[j] = sprite
	}
	return sprites, nil
}
