package asciify

import (
	"image"

	"github.com/disintegration/imaging"
)

var shade = []string{" ", "░", "▒", "▓", "█"}

// ToASCII Convert an image to ascii
func ToASCII(srcImg image.Image, width int, height int, brightness uint8) []string {
	resizedImg := imaging.Resize(srcImg, width, height, imaging.Lanczos)
	grayImg := imaging.Grayscale(resizedImg)
	grayImgBounds := grayImg.Bounds()
	result := make([]string, 0)
	for y := grayImgBounds.Min.Y; y < grayImgBounds.Max.Y; y++ {
		//row := []string{}
		for x := grayImgBounds.Min.X; x < grayImgBounds.Max.X; x++ {
			g, _, _, _ := grayImg.At(x, y).RGBA()
			result = append(result, shade[uint8(g)/brightness])
		}
		result = append(result, "\n")
	}
	return result
}
