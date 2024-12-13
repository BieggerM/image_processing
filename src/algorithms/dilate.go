package algorithms

import (
	"fmt"
	"image"
	"time"
	"github.com/BieggerM/image_processing_golang/util"
	"image/color"
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

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {	
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

	var (
		r uint32 
		g uint32
		b uint32
	)

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			r, g, b, _ = img.At(i, j).RGBA()
			if r == 255 && g == 255 && b == 255 {
				return color.RGBA{255, 255, 255, 255}
			}
		}
	}

	return color.RGBA{0, 0, 0, 255}
}