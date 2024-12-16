package algorithms

import (
	"fmt"
	"image"
	"image/color"
	"time"
	"sync"
	"github.com/BieggerM/image_processing_golang/util"
)
/*
Background_subtract is a function that subtracts the background of an image
It uses the substraction function to subtract the background of an image and measure the time it takes to do so

It times the following operations:
1. Open reference image
2. Open input image
3. Compatibility check
4. Background subtraction
5. Save image
*/

func Background_subtract(reference string, input string, threshold float64, hsv bool, multithreaded bool, numberofthreads int) {
	start := time.Now()
	fmt.Println("-----Reading Reference Image-----")
	
	// open reference image
	refImg, err := util.LoadImage(reference)
	if err != nil {
		fmt.Println("Failed to load reference image: ", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/bgsubtraction) finish open reference image \n", elapsed.Seconds())
	
	// open input image
	inputImg, err := util.LoadImage(input)
	if err != nil {
		fmt.Println("Failed to load input image: ", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/bgsubtraction) finish open input image \n", elapsed.Seconds())

	// check if dimensions of the images are the same
	err = util.CheckCompatibility(refImg, inputImg)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/bgsubtraction) finish compatibility check \n", elapsed.Seconds())

	// create output image
	outputImg := image.NewRGBA(refImg.Bounds())

	bounds := refImg.Bounds()
	if multithreaded {
		// use multithreaded version with syncgroups that ensure all threads are done before continuing
		var wg sync.WaitGroup
		rowsPerWorker := (bounds.Max.Y - bounds.Min.Y) / numberofthreads

		for w := 0; w < numberofthreads; w++ {
			wg.Add(1)
			go func(worker int) {
				defer wg.Done()
				// calculate the start and end of the rows for each worker
				startY := bounds.Min.Y + worker*rowsPerWorker
				endY := startY + rowsPerWorker
				// if it is the last worker, set the end to the last row
				if worker == numberofthreads-1 {
					endY = bounds.Max.Y
				}
				// call the substraction function for the rows
				substraction(startY, endY, bounds.Min.X, bounds.Max.X, hsv, refImg, inputImg, threshold, outputImg)
			}(w)
		}
		wg.Wait()
	} else {
		// call the substraction function for the whole image
		substraction(bounds.Min.Y, bounds.Max.Y, bounds.Min.X, bounds.Max.X, hsv, refImg, inputImg, threshold, outputImg)
	}
	
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/bgsubtraction) finish background subtraction \n", elapsed.Seconds())

	// save the output image
	err = util.SaveImage("../out/output.jpg", outputImg)
	if err != nil {
		fmt.Println("Failed to save output image: ", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/bgsubtraction) finish save image in %s\n", elapsed.Seconds(), "../out/output.jpg")
}

/* 
substraction is the algorithm that subtracts the background of an image
It uses the ColorDifferenceRGB function from the util package to calculate the difference between two colors
It uses the RgbToHsv function from the util package to convert RGB colors to HSV colors
It uses the Weighted_hsv_distance function from the util package to calculate the difference between two HSV colors
*/
func substraction(startY int, endY int, minX int, maxX int, hsv bool, refImg image.Image, inputImg image.Image, threshold float64, outputImg *image.RGBA) {
	var diff float64
	// iterate over the image
	for y := startY; y < endY; y++ {
		for x := minX; x < maxX; x++ {
			if hsv {
				// convert RGB to HSV
				h1, s1, v1 := util.RgbToHsv(refImg.At(x, y))
				h2, s2, v2 := util.RgbToHsv(inputImg.At(x, y))
				// calculate the difference between the two colors
				diff = util.Weighted_hsv_distance(h1, s1, v1, h2, s2, v2, 1.0, 1.0, 1.0)
			} else {
				r, g, b, _ := refImg.At(x, y).RGBA()
				r1, g1, b1, _ := inputImg.At(x, y).RGBA()
				// calculate the difference between the two colors
				diff = util.ColorDifferenceRGB(color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}, color.RGBA{uint8(r1 >> 8), uint8(g1 >> 8), uint8(b1 >> 8), 255})
			}
			if diff < threshold {
				// set the pixel to black
				outputImg.Set(x, y, color.RGBA{0, 0, 0, 255})
			} else {
				// set the pixel to white
				outputImg.Set(x, y, color.RGBA{255, 255, 255, 255})
			}
		}
	}
}

