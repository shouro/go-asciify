package asciify

import (
	"image"
	"math"

	"github.com/disintegration/imaging"
)

//var shade = []string{" ", "░", "▒", "▓", "█"}
var shade = []string{"█", "▓", "▒", "░", " "}

const uint8Max = 256

type line struct {
	n    int
	data string
}

// ToASCII Convert an image to ascii
func ToASCII(srcImg image.Image, darkness uint8) []string {
	grayImg := imaging.Grayscale(srcImg)
	grayImgBounds := grayImg.Bounds()
	darkness = darkness + uint8(math.Ceil(float64(uint8Max)/float64(len(shade))))
	result := []string{}
	ch := make(chan line)
	for y := grayImgBounds.Min.Y; y < grayImgBounds.Max.Y; y++ {
		go func(y int) {
			row := ""
			for x := grayImgBounds.Min.X; x < grayImgBounds.Max.X; x++ {
				g, _, _, _ := grayImg.At(x, y).RGBA()
				row = row + shade[uint8(g)/darkness]
			}
			ch <- line{n: y, data: row}
		}(y)
	}
	m := make(map[int]string)
	for y := grayImgBounds.Min.Y; y < grayImgBounds.Max.Y; y++ {
		select {
		case ln := <-ch:
			m[ln.n] = ln.data
		}
	}
	for i := 0; i < len(m); i++ {
		result = append(result, m[i])
	}
	return result
}
