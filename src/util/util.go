package util

import (
	"image"
	"image/jpeg"
	"os"
	"math"
	"image/color"
	"fmt"
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

/*
checkCompatibility is a function that checks if two images are compatible
It returns an error if the images are not compatible
*/
func CheckCompatibility(refImg image.Image, inputImg image.Image) error {
	if refImg.Bounds().Dx() != inputImg.Bounds().Dx() || refImg.Bounds().Dy() != inputImg.Bounds().Dy() {
		return fmt.Errorf("images are not compatible")
	}
	return nil
}

func ColorDifferenceRGB(c1, c2 color.RGBA) float64 {
	rDiff := float64(c1.R) - float64(c2.R)
	gDiff := float64(c1.G) - float64(c2.G)
	bDiff := float64(c1.B) - float64(c2.B)
	return math.Abs(rDiff + gDiff + bDiff) / 3.0
}
