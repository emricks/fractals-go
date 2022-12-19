package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

const (
	imageHeight = 18000
	iterations  = 100000000
	vertexAngle = 120
	hexRadians  = (180 - vertexAngle) / 2 * math.Pi / 180
)

func main() {
	nonHexHorizontal := math.Tan(hexRadians)*(imageHeight/2) + 1
	imageWidth := nonHexHorizontal * 4
	rectangle := image.Rect(0, 0, int(imageWidth), imageHeight)
	img := image.NewGray(rectangle)
	vertices := []complex128{
		complex(nonHexHorizontal, 0),
		complex(imageWidth-nonHexHorizontal, 0),
		complex(imageWidth, imageHeight/2),
		complex(imageWidth-nonHexHorizontal, imageHeight),
		complex(nonHexHorizontal, imageHeight),
		complex(0, imageHeight/2),
	}
	var current = complex(imageWidth, imageHeight)
	for i := 0; i < iterations; i++ {
		dice := rand.Int() % 6
		current = (vertices[dice]*2 + current) / 3
		img.Set(int(real(current)), int(imag(current)), color.White)
	}

	fmt.Println("Saving image")
	saveImage("SerpinskiHexagon.png", img)
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
