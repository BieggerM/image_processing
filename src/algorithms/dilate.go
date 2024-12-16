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
Dilate is a function that dilates an image
It uses the dilatePixel function to dilate an image and measure the time it takes to do so

It times the following operations:
1. Open image
2. Create output image
3. Dilate
4. Save image
*/

func Dilate(input string, radius int, multithreaded bool, numberofthreads int)  {
	start := time.Now()
	fmt.Println("-----Reading Image-----")

	// open image
	img, err := util.LoadImage(input)
	if err != nil {
		fmt.Println("Failed to load image: ", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/dilate) finish open image \n", elapsed.Seconds())

	// create output image
	outputImg := image.NewRGBA(img.Bounds())
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/dilate) finish create output image \n", elapsed.Seconds())

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
				for y := startY; y < endY; y++ {
					for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
						// dilate the pixel
						outputImg.Set(x, y, dilatePixel(img, x, y, radius))
					}
				}
			}(w)
		}
		wg.Wait()

	} else {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				// dilate the pixel
				outputImg.Set(x, y, dilatePixel(img, x, y, radius))
			}
		}
	}

	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/dilate) finish dilate \n", elapsed.Seconds())

	// save image
	err = util.SaveImage("../out/dilate.jpg", outputImg)
	if err != nil {
		fmt.Println("Failed to save image: ", err)
		return
	}

	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/dilate) finish save image ../out/dilate.jpg\n", elapsed.Seconds())
}	

func dilatePixel(img image.Image, x, y, radius int) color.Color {
	// get the min and max values for the x and y coordinates
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

	// iterate the range of pixels
	for i := minX; i <= maxX; i++ {
        for j := minY; j <= maxY; j++ {
			// look at the pixel at the current coordinates
            r, g, b, _ := img.At(i, j).RGBA()
			// if any of the RGB values are white, return white
            if r == 65535 && g == 65535 && b == 65535 { // Check if the pixel is white
                return color.RGBA{255, 255, 255, 255} // Set the current pixel to white
            }
        }
    }
	// if no white pixels are found, return black
	return color.RGBA{0, 0, 0, 255}
}