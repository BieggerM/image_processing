package algorithms

import (
	"fmt"
	"image"
	"time"
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

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {	
			outputImg.Set(x, y, img.At(x, y))
			for i := -radius; i <= radius; i++ {
				for j := -radius; j <= radius; j++ {
					if x+i >= bounds.Min.X && x+i < bounds.Max.X && y+j >= bounds.Min.Y && y+j < bounds.Max.Y {
						outputImg.Set(x, y, img.At(x+i, y+j))
					}
				}
			}
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

