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
var dark int
var original bool
var darkMin = 0
var darkMax = 100

func init() {
	flag.IntVar(&width, "width", 80, "New width to resize")
	flag.IntVar(&width, "x", 80, "New width to resize (shorthand)")

	flag.IntVar(&height, "height", 0, "New height to resize")
	flag.IntVar(&height, "y", 0, "New height to resize (shorthand)")

	flag.IntVar(&dark, "dark", 0, "darkness")
	flag.IntVar(&dark, "d", 0, "darkness (shorthand)")

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

	if dark < 0 || dark > 100 {
		fmt.Println("darkness must be between 0 and 100")
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
	resizedImg := imaging.Resize(srcImg, width, height, imaging.Lanczos)
	result := asciify.ToASCII(resizedImg, uint8(dark))
	stdoutPrinter(result)
}

func stdoutPrinter(result []string) {
	for _, row := range result {
		fmt.Print(row + "\n")
	}
}
