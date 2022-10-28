/*
	Package image defines the Image interface:

	package image

	type Image interface {
		ColorModel() color.Model
		Bounds() Rectangle // Is actually an image.Rectangle, as the declaration is inside package image
		At(x, y int) color.Color
	}
*/

package goTour04

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8((x + y) / 2), uint8(x * y), uint8(x ^ y), 255}
}

func RunExercise25() {
	m := Image{600, 300}
	pic.ShowImage(m)
}
