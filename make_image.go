package main

import (
	"image"
	"golang.org/x/tour/pic"
	"image/color"
)

type Image struct {
	x, y int
	r, g, b, a uint8
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.x, i.y)

}

func (i *Image) At (x, y int) color.Color {
	return color.RGBA{i.r, i.g, i.b, i.a}
}

func main() {
	m := &Image{
		x: 100,
		y: 200,
		r: 255,
		g: 128,
		b: 128,
		a: 255,
	}
	pic.ShowImage(m)
}