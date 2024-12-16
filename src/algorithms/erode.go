package algorithms

import (
	"fmt"
	"image"
	"time"
	"image/color"
	"github.com/BieggerM/image_processing_golang/util"
	"sync"
)

/*
Erode is a function that erodes an image
It uses the erodePixel function to erode an image and measure the time it takes to do so

It times the following operations:
1. Open image
2. Create output image
3. Erode
4. Save image
*/

func Erode(input string, radius int, multithreaded bool, numberofthreads int) {
	start := time.Now()

	// open image
	img, err := util.LoadImage(input)
	if err != nil {
		fmt.Println("Failed to load image: ", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/erode) finish open image \n", elapsed.Seconds())

	// create output image
	outputImg := image.NewRGBA(img.Bounds())
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/erode) finish create output image \n", elapsed.Seconds())

	if multithreaded {
		// use multithreaded version with syncgroups that ensure all threads are done before continuing
		var wg sync.WaitGroup
		rowsPerWorker := (img.Bounds().Max.Y - img.Bounds().Min.Y) / numberofthreads
		for w := 0; w < numberofthreads; w++ {
			wg.Add(1)
			go func(worker int) {
				defer wg.Done()
				// calculate the start and end of the rows for each worker
				startY := img.Bounds().Min.Y + worker*rowsPerWorker
				endY := startY + rowsPerWorker
				// if it is the last worker, make sure to include the remaining rows
				if worker == numberofthreads-1 {
					endY = img.Bounds().Max.Y
				}
				// erode the workers rows
				for y := startY; y < endY; y++ {
					for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
						outputImg.Set(x, y, erodePixel(img, x, y, radius))
					}
				}
			}(w)
		}
		wg.Wait()
	} else {
		// erode the image
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				outputImg.Set(x, y, erodePixel(img, x, y, radius))
			}
		}
	}

	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/erode) finish erode \n", elapsed.Seconds())

	// save the image
	err = util.SaveImage("../out/erode.jpg", outputImg)
	if err != nil {
		fmt.Println("Failed to save image: ", err)
		return
	}

	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/erode) finish save image ../out/erode.jpg\n", elapsed.Seconds())
}	

func erodePixel(img image.Image, x, y, radius int) color.Color {
	// check the pixels around the current pixel
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

	// iterate over the pixels around the current pixel
	for i := minX; i <= maxX; i++ {
        for j := minY; j <= maxY; j++ {
			// check if the pixel is black
            r, g, b, _ := img.At(i, j).RGBA()
            if r == 0 && g == 0 && b == 0 { // Check if the pixel is black
                return color.RGBA{0, 0, 0, 255} // Set the current pixel to black
            }
        }
    }

	// if no black pixels are found, return white
	return color.RGBA{255, 255, 255, 255}
}