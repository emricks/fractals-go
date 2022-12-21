package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

const (
	imageSize  = 10000
	iterations = 100000000
)

func main() {
	rectangle := image.Rect(0, 0, imageSize, imageSize)
	img := image.NewGray(rectangle)
	var pt0 = complex(imageSize*0.5, 0)
	var pt1 = complex(imageSize*0.8, imageSize*0.2)
	var pt2 = complex(imageSize*0.8, imageSize*0.6)
	var pt3 = complex(imageSize*0.5, imageSize*0.8)
	var pt4 = complex(imageSize*0.2, imageSize*0.6)
	var pt5 = complex(imageSize*0.2, imageSize*0.2)
	var current = complex(0, imageSize/2)
	var current2 = complex(imageSize/2, 0)
	for i := 0; i < iterations; i++ {
		dice := rand.Int() % 3
		if dice == 0 {
			current = (current + pt1) / 2
			current2 = (current2 + pt0) / 2
		}
		if dice == 1 {
			current = (current + pt3) / 2
			current2 = (current2 + pt2) / 2
		}
		if dice == 2 {
			current = (current + pt5) / 2
			current2 = (current2 + pt4) / 2
		}
		img.Set(int(real(current)), int(imag(current)), color.White)
		img.Set(int(real(current2)), int(imag(current2)), color.White)
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
