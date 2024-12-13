package algorithms

import (
	"fmt"
	"image"
	"time"
	"image/color"
	"github.com/BieggerM/image_processing_golang/util"
)

func Dilate(input string, radius int)  {
	start := time.Now()

	fmt.Println("-----Reading Image-----")
	img, err := util.LoadImage(input)
	if err != nil {
		fmt.Println("Failed to load image: ", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("[%s] Image loaded in \n", elapsed)

	fmt.Println("-----Dilating Image-----")
	outputImg := image.NewRGBA(img.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			outputImg.Set(x, y, dilatePixel(img, x, y, radius))
		}
	}

	elapsed = time.Since(start)
	fmt.Printf("[%s] Image dilated in \n", elapsed)

	fmt.Println("-----Saving Image-----")
	err = util.SaveImage("../out/dilate.jpg", outputImg)
	if err != nil {
		fmt.Println("Failed to save image: ", err)
		return
	}

	elapsed = time.Since(start)
	fmt.Printf("[%s] Image saved in \n", elapsed)
}	

func dilatePixel(img image.Image, x, y, radius int) color.Color {
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

	for i := minX; i <= maxX; i++ {
        for j := minY; j <= maxY; j++ {
            r, g, b, _ := img.At(i, j).RGBA()
            if r == 65535 && g == 65535 && b == 65535 { // Check if the pixel is white
                return color.RGBA{255, 255, 255, 255} // Set the current pixel to white
            }
        }
    }

	return color.RGBA{0, 0, 0, 255}
}