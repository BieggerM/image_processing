package util

import (
	"image"
	"image/jpeg"
	"image/color"
	"os"
)

func LoadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func SaveImage(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}
	return nil
}

func ErodePixel(img image.Image, x, y, radius int) color.Color {
	minX := x - radius
	if minX < img.Bounds().Min.X {
		minX = img.Bounds().Min.X
	}
	minY := y - radius
	if minY < img.Bounds().Min.Y {
		minY = img.Bounds().Min.Y
	}
	maxX := x + radius
	if maxX > img.Bounds().Max.X {
		maxX = img.Bounds().Max.X
	}
	maxY := y + radius
	if maxY > img.Bounds().Max.Y {
		maxY = img.Bounds().Max.Y
	}

	var r, g, b, a uint32
	r = 255
	g = 255
	b = 255
	a = 255

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			r_, g_, b_, a_ := img.At(i, j).RGBA()
			if r_ < r {
				r = r_
			}
			if g_ < g {
				g = g_
			}
			if b_ < b {
				b = b_
			}
			if a_ < a {
				a = a_
			}
		}
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}