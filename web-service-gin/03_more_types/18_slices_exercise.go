package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picToReturn := make([][]uint8, dx)

	for x := 0; x < dx; x++ {
		picToReturn[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			picToReturn[x][y] = uint8((x + y) / 2)
		}
	}

	return picToReturn
}

func main() {
	pic.Show(Pic)
}
