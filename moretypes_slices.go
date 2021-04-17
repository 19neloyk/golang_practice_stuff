package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dx)
	for indx := range image{
		cur_dim := make ([]uint8,dy)
		for indy := range cur_dim {
			cur_dim[indy] = uint8((4*indx + 4*indy)/2)
		}
		image[indx] = cur_dim
	}
	
	return image
}

func main() {
	pic.Show(Pic)
}
