package level

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/faiface/pixel"
	"github.com/svera/quarter"
	"github.com/svera/quarter/collision"
)

// Returned errors
const (
	ErrorWrongFileFormat       = "Loaded levels file is not a valid JSON"
	ErrorVersionNotSupported   = "Version \"%s\" not supported"
	ErrorNoLevels              = "Levels file must have at least one level declared, none found"
	ErrorBoundTypeNotSupported = "Bound type \"%s\" is not supported"
	ErrorWrongBoundValues      = "Wrong or missing bound values"
)

// LevelsFile is the serialized form of Levels
type LevelsFile struct {
	Version string
	Levels  []struct {
		Name   string
		Limits struct {
			Min pixel.Vec
			Max pixel.Vec
		}
		Layers []struct {
			Name  string
			Image struct {
				Path string
			}
			Grid struct {
				Assets GridAssets
				Tiles  []struct {
					Asset   int
					X       float64
					Y       float64
					Bounded bool
					Extra   json.RawMessage
				}
				Extra json.RawMessage
			}
			Bounds []struct {
				Type   string
				Values json.RawMessage
			}
			Extra json.RawMessage
		}
		Extra json.RawMessage
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

type BoundRectValues struct {
	Min pixel.Vec
	Max pixel.Vec
}

// Load validates a levels file and returns its information as a []Level
func Load(r io.Reader) ([]Level, error) {
	data := LevelsFile{}
	levels := make([]Level, len(data.Levels))

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

	for _, currentLevel := range data.Levels {
		level := Level{Limits: pixel.R(
			currentLevel.Limits.Min.X,
			currentLevel.Limits.Min.Y,
			currentLevel.Limits.Max.X,
			currentLevel.Limits.Max.Y,
		)}
		for i, currentLayer := range currentLevel.Layers {
			level.Layers = append(level.Layers, Layer{})
			if path := strings.TrimSpace(currentLayer.Image.Path); path != "" {
				img, err := quarter.LoadPicture(path)
				if err != nil {
					return nil, err
				}
				level.Layers[i].image = pixel.NewSprite(img, img.Bounds())
			}
			boundedTiles := make([]collision.Shaper, 0, len(currentLayer.Grid.Tiles))
			boundShapes := make([]collision.Shaper, 0, len(currentLayer.Bounds))
			if len(currentLayer.Grid.Tiles) != 0 {
				level.Layers[i].Grid = &Grid{
					tileWidth:  currentLayer.Grid.Assets.Width,
					tileHeight: currentLayer.Grid.Assets.Height,
				}
				level.Layers[i].Grid.Assets, err = loadGridAssets(currentLayer.Grid.Assets)
				if err != nil {
					return nil, err
				}
				level.Layers[i].Grid.Tiles = make([]GridTile, len(currentLayer.Grid.Tiles))
				for k, val := range currentLayer.Grid.Tiles {
					tl := GridTile{
						Asset:  val.Asset,
						Coords: pixel.V(val.X, val.Y),
					}
					level.Layers[i].Grid.Tiles[k] = tl
					if val.Bounded {
						boundedTiles = append(boundedTiles, level.Layers[i].Grid.TileBoundingBox(pixel.V(val.X, val.Y)))
					}
				}
			}
			if len(currentLayer.Bounds) > 0 {
				for _, bound := range currentLayer.Bounds {
					switch bound.Type {
					case "box":
						values := BoundRectValues{}
						err := json.Unmarshal(bound.Values, &values)
						if err != nil {
							return nil, fmt.Errorf(ErrorWrongBoundValues)
						}
						boundShapes = append(boundShapes, collision.NewBoundingBox(pixel.V(values.Min.X, values.Min.Y), pixel.V(values.Max.X, values.Max.Y)))
					default:
						return nil, fmt.Errorf(ErrorBoundTypeNotSupported, bound.Type)
					}
				}
			}
			level.Layers[i].Bounds = make([]collision.Shaper, 0, len(boundShapes)+len(boundedTiles))
			level.Layers[i].Bounds = append(level.Layers[i].Bounds, boundShapes...)
			level.Layers[i].Bounds = append(level.Layers[i].Bounds, boundedTiles...)
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
				assets.Offset.Y,
				x+assets.Width,
				assets.Offset.Y+assets.Height,
			),
		)
		sprites[j] = sprite
	}
	return sprites, nil
}
