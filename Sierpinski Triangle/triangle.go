package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

const (
	imageSize  = 15000
	iterations = 200000000
)

func main() {
	rectangle := image.Rect(0, 0, imageSize, imageSize)
	img := image.NewGray(rectangle)
	var pt0 = complex(imageSize/2, 0)
	var pt1 = complex(0, imageSize)
	var pt2 = complex(imageSize, imageSize)
	var current = complex(0, imageSize/2)
	for i := 0; i < iterations; i++ {
		dice := rand.Int() % 3
		if dice == 0 {
			current = (current + pt0) / 2
		}
		if dice == 1 {
			current = (current + pt1) / 2
		}
		if dice == 2 {
			current = (current + pt2) / 2
		}
		img.Set(int(real(current)), int(imag(current)), color.White)
	}
	saveImage("SerpinskiTriangle.png", img)
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
