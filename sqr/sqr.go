package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

const (
	imageSize  = 15000
	iterations = 1000000000
)

func main() {
	start := time.Now()
	rectangle := image.Rect(0, 0, imageSize, imageSize)
	img := image.NewGray(rectangle)
	vertices := []complex128{
		complex(0, 0),
		complex(imageSize/2, 0),
		complex(imageSize, 0),
		complex(imageSize, imageSize/2),
		complex(imageSize, imageSize),
		complex(imageSize/2, imageSize),
		complex(0, imageSize),
		complex(0, imageSize/2),
	}
	var current = complex(imageSize/2, imageSize/2)
	for i := 0; i < iterations; i++ {
		dice := rand.Int() % 8
		current = (vertices[dice]*2 + current) / 3
		img.Set(int(real(current)), int(imag(current)), color.White)
	}
	end := time.Now()
	fmt.Printf("Calculations done in %v seconds, Saving image", end.Sub(start).Seconds())
	saveImage("SerpinskiCarpet.png", img)
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
