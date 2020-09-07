package main

import (
	"bytes"
	"image/color"
	"io"
	"os"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/svera/quarter/bound"
	"github.com/svera/quarter/level"
)

type Level struct {
	levels map[string]level.Level
	Bounds map[string][]bound.Shaper
	circle *bound.Circle
}

func NewLevel(dataFile string) (*Level, error) {
	var buf bytes.Buffer
	data, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	tee := io.TeeReader(data, &buf)
	lvl, err := level.Deserialize(tee)
	if err != nil {
		return nil, err
	}
	levelBounds, err := bound.Deserialize(&buf)
	if err != nil {
		return nil, err
	}

	return &Level{
		Bounds: levelBounds,
		levels: lvl,
		circle: bound.NewCircle(0, 40, 18),
	}, nil
}

func (l *Level) Draw(target *pixelgl.Canvas, debug *color.RGBA, imd *imdraw.IMDraw) {
	l.levels["level1"].Draw(target)
	if debug != nil {
		for _, b := range l.Bounds["level1-background"] {
			b.Shape().Draw(debug, imd, target)
		}
	}

	l.circle.Draw(&color.RGBA{0, 255, 0, 16}, imd, target)
}
