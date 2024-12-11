package algorithms

import (
	"fmt"
	"image"
	"time"
	"image/color"
	"github.com/BieggerM/image_processing_golang/util"
)

func Erode(input string, radius int)  {
	start := time.Now()

	fmt.Println("-----Reading Image-----")
	img, err := util.LoadImage(input)
	if err != nil {
		fmt.Println("Failed to load image: ", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("[%s] Image loaded in \n", elapsed)

	fmt.Println("-----Eroding Image-----")
	outputImg := image.NewRGBA(img.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			outputImg.Set(x, y, erodePixel(img, x, y, radius))
		}
	}
	
	elapsed = time.Since(start)
	fmt.Printf("[%s] Image eroded in \n", elapsed)

	fmt.Println("-----Saving Image-----")
	err = util.SaveImage("../out/erode.jpg", outputImg)
	if err != nil {
		fmt.Println("Failed to save image: ", err)
		return
	}

	elapsed = time.Since(start)
	fmt.Printf("[%s] Image saved in \n", elapsed)
	

}

func erodePixel(img image.Image, x, y, radius int) color.Color {
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