package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
	"time"
)

const (
	startpoint  = complex(-.7, .7)
	endx        = -.3
	imageWidth  = 12000
	imageHeight = 8000
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	start := time.Now()
	renderImage(img, startpoint, endx)
	end := time.Now()
	fmt.Printf("Calculations done in %v seconds, Saving image", end.Sub(start).Seconds())
	saveImage("mandleset.png", img)
}

func mandlebrot(coordinate complex128, limit int) int {
	var z = complex(0, 0)
	for i := 0; i < limit; i++ {
		z = z*z + coordinate
		if cmplx.Abs(z) > 2 {
			return i
		}
	}
	return limit
}

func renderImage(img *image.RGBA, startpoint complex128, endx float64) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	realx := endx - real(startpoint)
	xdiff := math.Abs(realx)
	pixelsize := xdiff / float64(width)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tempr := real(startpoint) + float64(x)*pixelsize
			tempi := imag(startpoint) - float64(y)*pixelsize
			coord := complex(tempr, tempi)
			i := 255 - mandlebrot(coord, 255)
			shade := color.RGBA{uint8(i), uint8(i), uint8(i), 255}
			img.Set(x, y, shade)
		}
	}
}

func saveImage(fileName string, img image.Image) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
