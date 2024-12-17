package util

import (
	"image"
	"image/jpeg"
	"os"
)

/*
LoadImage is a function that loads an image from a file
*/
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

/*
SaveImage is a function that saves an image to a file
*/
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

