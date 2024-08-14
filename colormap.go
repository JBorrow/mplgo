package mplgo

import (
	"image"
	"image/color"
	"math"
)

var BAD_COLOR = color.RGBA{1.0, 1.0, 1.0, 0.0}

type ColorMap struct {
	data   []color.RGBA
	name   string
	steps  int
	fsteps float64
}

func (m ColorMap) Map(in float64) color.RGBA {
	in = math.Max(in, 0.0)
	in = math.Min(in, 1.0)
	
	rounded := math.Round(in * (m.fsteps - 1.0))

	if math.IsNaN(rounded) {
		return BAD_COLOR
	}

	idx := int(math.Round(in * (m.fsteps - 1.0)))

	return m.data[idx]
}

func (m ColorMap) MapArray(in [][]float64) [][]color.RGBA {
	output := make([][]color.RGBA, len(in))

	for i, line := range in {
		outputLine := make([]color.RGBA, len(line))
		for j, val := range line {
			outputLine[j] = m.Map(val)
		}
		output[i] = outputLine
	}

	return output
}

func (m ColorMap) MapArrayToImage(in [][]float64) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, len(in), len(in[0])))

	for i, line := range in {
		for j, val := range line {
			img.Set(j, i, m.Map(val))
		}
	}

	return img
}
