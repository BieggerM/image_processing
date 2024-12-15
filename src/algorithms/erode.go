package algorithms

import (
	"fmt"
	"image"
	"time"
	"image/color"
	"github.com/BieggerM/image_processing_golang/util"
	"sync"
)

func Erode(input string, radius int, multithreaded bool, numberofthreads int) {
	start := time.Now()

	img, err := util.LoadImage(input)
	if err != nil {
		fmt.Println("Failed to load image: ", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/erode) finish open image \n", elapsed.Seconds())

	outputImg := image.NewRGBA(img.Bounds())
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/erode) finish create output image \n", elapsed.Seconds())

	if multithreaded {
		var wg sync.WaitGroup
		rowsPerWorker := (img.Bounds().Max.Y - img.Bounds().Min.Y) / numberofthreads
		for w := 0; w < numberofthreads; w++ {
			wg.Add(1)
			go func(worker int) {
				defer wg.Done()
				startY := img.Bounds().Min.Y + worker*rowsPerWorker
				endY := startY + rowsPerWorker
				if worker == numberofthreads-1 {
					endY = img.Bounds().Max.Y
				}
				for y := startY; y < endY; y++ {
					for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
						outputImg.Set(x, y, erodePixel(img, x, y, radius))
					}
				}
			}(w)
		}
		wg.Wait()
	} else {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				outputImg.Set(x, y, erodePixel(img, x, y, radius))
			}
		}
	}

	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/erode) finish erode \n", elapsed.Seconds())

	err = util.SaveImage("../out/erode.jpg", outputImg)
	if err != nil {
		fmt.Println("Failed to save image: ", err)
		return
	}

	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/erode) finish save image ../out/erode.jpg\n", elapsed.Seconds())
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

	for i := minX; i <= maxX; i++ {
        for j := minY; j <= maxY; j++ {
            r, g, b, _ := img.At(i, j).RGBA()
            if r == 0 && g == 0 && b == 0 { // Check if the pixel is black
                return color.RGBA{0, 0, 0, 255} // Set the current pixel to black
            }
        }
    }

	return color.RGBA{255, 255, 255, 255}
}