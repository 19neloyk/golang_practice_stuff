package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

//Add the three methods under this to become an image
type Image struct{
	dim_x int
	dim_y int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.dim_x, i.dim_y)
}

func (i Image) At(x int ,y int) color.Color {
	v := uint8((x*y^(1/2))/2)
	return color.RGBA{v, v, 255, 255} 
}

func main() {
	m := Image{dim_x:500,dim_y:500}
	pic.ShowImage(m)
}
