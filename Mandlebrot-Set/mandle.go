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
	"sync"
	"time"
)

const (
	startpoint   = complex(-2, 1)
	endpointreal = 1
	//startpoint   = complex(-.4, .665)
	//endpointreal = -.3
	imageWidth  = 6000
	imageHeight = 4000
	iterations  = 1024
	cpus        = 10
)

var (
	colors = palette.Rainbow(iterations+1, palette.Hue(.6), palette.Hue(0), .7, .7, 1)
	//colors = palette.Heat(iterations+1, .5)
	//colors = palette.Radial(iterations+1, palette.Blue, palette.Red, .7)
	//colors = palette.Rainbow(iterations+1, palette.Hue(0.7), palette.Hue(0), .7, .7, 1)
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(cpus)
	for cpu := 0; cpu < cpus; cpu++ {
		assignedCPU := cpu
		go func() {
			renderPerCPU(cpus, assignedCPU, img, startpoint, endpointreal)
			wg.Done()
		}()
	}
	wg.Wait()
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

// render a portion of the image according to the CPU assigned
func renderPerCPU(cpuCount int, assignedCpu int, img *image.RGBA, startcoord complex128, endX float64) {
	subHeight := img.Bounds().Dy() / cpuCount
	subY := subHeight * assignedCpu
	pixelsize := getPixelSize(img, startcoord, endX)
	substartpoint := pixelCoordToMandelCoord(0, subY, pixelsize, startcoord)
	subRectangle := image.Rect(0, subY, img.Bounds().Dx(), subY+subHeight)
	subImg := img.SubImage(subRectangle).(*image.RGBA)
	renderImage(subImg, substartpoint, endX)
}

func renderImage(img *image.RGBA, startcoord complex128, endx float64) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	pixelsize := getPixelSize(img, startcoord, endx)
	fmt.Printf("Calculating segment width: %d, height %d, at %v\n", width, height, startcoord)
	minpoint := img.Bounds().Min

	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			coord := pixelCoordToMandelCoord(x, y, pixelsize, startcoord)
			i := mandlebrot(coord, iterations)
			shade := iter2Palette(i)
			currentPoint := minpoint.Add(image.Point{X: x, Y: y})
			img.Set(currentPoint.X, currentPoint.Y, shade)
		}
	}
}

func getPixelSize(img image.Image, startpoint complex128, endx float64) float64 {
	realx := endx - real(startpoint)
	xdiff := math.Abs(realx)
	return xdiff / float64(img.Bounds().Dx())
}

func pixelCoordToMandelCoord(x, y int, pixelsize float64, startpoint complex128) complex128 {
	tempr := real(startpoint) + float64(x)*pixelsize
	tempi := imag(startpoint) - float64(y)*pixelsize
	return complex(tempr, tempi)
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
