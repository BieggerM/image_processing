package algorithms

import (
	"fmt"
	"image"
	"time"
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
			outputImg.Set(x, y, util.ErodePixel(img, x, y, radius))
		}
	}
	
	elapsed = time.Since(start)
	fmt.Printf("[%s] Image eroded in \n", elapsed)

	fmt.Println("-----Saving Image-----")
	err = util.SaveImage("output.png", outputImg)
	if err != nil {
		fmt.Println("Failed to save image: ", err)
		return
	}

	elapsed = time.Since(start)
	fmt.Printf("[%s] Image saved in \n", elapsed)
	

}