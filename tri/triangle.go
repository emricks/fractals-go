package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

const (
	imageSize  = 1000
	iterations = 100000000
)

func main() {
	rectangle := image.Rect(0, 0, imageSize, imageSize)
	img := image.NewGray(rectangle)
	var pt0 = complex(0, imageSize)
	var pt1 = complex(666, 0)
	var pt2 = complex(imageSize, imageSize/2)
	var pt3 = complex(imageSize, 0)
	var pt4 = complex(0, 0)
	var pt5 = complex(imageSize, imageSize)
	var current = complex(0, imageSize/2)
	var current2 = complex(imageSize, 1)
	var current3 = complex(0, 1)
	var current4 = complex(imageSize, imageSize-1)
	for i := 0; i < iterations; i++ {
		dice := rand.Int() % 3
		if dice == 0 {
			current = (current + pt0) / 2
			current2 = (current2 + pt1) / 2
			current3 = (current3 + pt0) / 2
			current4 = (current4 + pt0) / 2
		}
		if dice == 1 {
			current = (current + pt1) / 2
			current2 = (current2 + pt2) / 2
			current3 = (current3 + pt1) / 2
			current4 = (current4 + pt2) / 2
		}
		if dice == 2 {
			current = (current + pt2) / 2
			current2 = (current2 + pt3) / 2
			current3 = (current3 + pt4) / 2
			current4 = (current4 + pt5) / 2
		}
		img.Set(int(real(current)), int(imag(current)), color.White)
		img.Set(int(real(current2)), int(imag(current2)), color.White)
		img.Set(int(real(current3)), int(imag(current3)), color.White)
		img.Set(int(real(current4)), int(imag(current4)), color.White)
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
