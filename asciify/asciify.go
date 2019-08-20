package asciify

import (
	"fmt"
	"image"
	"math"

	"github.com/disintegration/imaging"
)

var shade = []string{" ", "░", "▒", "▓", "█"}

const uint8Max = 256

// ToASCII Convert an image to ascii
func ToASCII(srcImg image.Image, width int, height int, darkness uint8) [][]string {
	resizedImg := imaging.Resize(srcImg, width, height, imaging.Lanczos)
	grayImg := imaging.Grayscale(resizedImg)
	grayImgBounds := grayImg.Bounds()
	darkness = darkness + uint8(math.Ceil(float64(uint8Max)/float64(len(shade))))
	result := [][]string{}
	for y := grayImgBounds.Min.Y; y < grayImgBounds.Max.Y; y++ {
		row := []string{}
		for x := grayImgBounds.Min.X; x < grayImgBounds.Max.X; x++ {
			g, _, _, _ := grayImg.At(x, y).RGBA()
			fmt.Println(uint8(g))
			row = append(row, shade[uint8(g)/darkness])
		}
		result = append(result, row)
	}
	return result
}
