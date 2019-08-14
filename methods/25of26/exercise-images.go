package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width, Height int
}

func (image Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (image Image) At(x, y int) color.Color {
	return color.RGBA{R: 0x5e, G: 0xc9, B: 0xc4, A: 0xff}
}

func main() {
	m := Image{
		Width:  300,
		Height: 250,
	}
	pic.ShowImage(m)
}
