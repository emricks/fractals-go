package main

import (
	"fmt"
	"gonum.org/v1/plot/palette"
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
	imageWidth  = 6000
	imageHeight = 4000
	iterations  = 512
)

var (
	colors = palette.Rainbow(iterations+1, palette.Blue, palette.Red, .5, .5, 1)
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
			i := mandlebrot(coord, iterations)
			shade := iter2Palette(i)
			img.Set(x, y, shade)
		}
	}
}

func iter2Green(i int) color.RGBA {
	return color.RGBA{R: uint8(i) / 100, G: uint8(i), B: uint8(i) / 100, A: 255}
}

func iter2Palette(i int) color.Color {
	return colors.Colors()[i]
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
