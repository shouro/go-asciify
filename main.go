package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/disintegration/imaging"
	"github.com/shouro/go-asciify/asciify"
)

var width int
var height int
var bright int
var original bool

func init() {
	flag.IntVar(&width, "width", 80, "New width to resize")
	flag.IntVar(&width, "x", 80, "New width to resize (shorthand)")

	flag.IntVar(&height, "height", 0, "New height to resize")
	flag.IntVar(&height, "y", 0, "New height to resize (shorthand)")

	flag.IntVar(&bright, "bright", 62, "Brightness")
	flag.IntVar(&bright, "b", 62, "Brightness (shorthand)")

	flag.BoolVar(&original, "original", false, "Using the image unmodified")
	flag.BoolVar(&original, "o", false, "Using the image unmodified (shorthand)")
}

func main() {
	flag.Parse()

	if width < 0 {
		fmt.Println("Nagative value as width is pointless")
		return
	}

	if height < 0 {
		fmt.Println("Nagative value as height is pointless")
		return
	}

	if bright < 52 || bright > 255 {
		fmt.Println("Brightness must be between 52 and 255")
		return
	}

	tail := flag.Args()
	if taillen := len(tail); taillen == 0 {
		fmt.Println("No image path provided")
		return
	}

	imgPath := tail[0]
	srcImg, err := imaging.Open(imgPath, imaging.AutoOrientation(true))
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	if original == true {
		width = srcImg.Bounds().Max.X
		height = srcImg.Bounds().Max.Y
	}

	result := asciify.ToASCII(srcImg, width, height, uint8(bright))
	stdoutPrinter(result)
}

func stdoutPrinter(result []string) {
	for _, c := range result {
		fmt.Print(c)
	}
}
