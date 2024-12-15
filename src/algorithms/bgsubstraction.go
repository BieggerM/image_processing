package algorithms

import (
	"fmt"
	"image"
	"image/color"
	"time"
	"sync"
	"github.com/BieggerM/image_processing_golang/util"
)

func Background_subtract(reference string, input string, threshold float64, hsv bool, multithreaded bool, numberofthreads int) {
	start := time.Now()
	fmt.Println("-----Reading Reference Image-----")
	refImg, err := util.LoadImage(reference)
	if err != nil {
		fmt.Println("Failed to load reference image: ", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/bgsubtraction) finish open reference image \n", elapsed.Seconds())

	inputImg, err := util.LoadImage(input)
	if err != nil {
		fmt.Println("Failed to load input image: ", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/bgsubtraction) finish open input image \n", elapsed.Seconds())

	err = checkCompatibility(refImg, inputImg)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/bgsubtraction) finish compatibility check \n", elapsed.Seconds())

	outputImg := image.NewRGBA(refImg.Bounds())

	bounds := refImg.Bounds()
	if multithreaded {
		var wg sync.WaitGroup
		rowsPerWorker := (bounds.Max.Y - bounds.Min.Y) / numberofthreads

		for w := 0; w < numberofthreads; w++ {
			wg.Add(1)
			go func(worker int) {
				defer wg.Done()
				startY := bounds.Min.Y + worker*rowsPerWorker
				endY := startY + rowsPerWorker
				if worker == numberofthreads-1 {
					endY = bounds.Max.Y
				}
				substraction(startY, endY, bounds.Min.X, bounds.Max.X, hsv, refImg, inputImg, threshold, outputImg)
			}(w)
		}
		wg.Wait()
	} else {
		substraction(bounds.Min.Y, bounds.Max.Y, bounds.Min.X, bounds.Max.X, hsv, refImg, inputImg, threshold, outputImg)
	}
	
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/bgsubtraction) finish background subtraction \n", elapsed.Seconds())

	err = util.SaveImage("../out/output.jpg", outputImg)
	if err != nil {
		fmt.Println("Failed to save output image: ", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("[%.7f] 0 (algorithm/bgsubtraction) finish save image in %s\n", elapsed.Seconds(), "../out/output.jpg")
}

func substraction(startY int, endY int, minX int, maxX int, hsv bool, refImg image.Image, inputImg image.Image, threshold float64, outputImg *image.RGBA) {
	stride := outputImg.Stride
	var diff float64
	for y := startY; y < endY; y++ {
		for x := minX; x < maxX; x++ {
			offset := y*stride + x*4
			if hsv {
				h1, s1, v1 := util.RgbToHsv(refImg.At(x, y))
				h2, s2, v2 := util.RgbToHsv(inputImg.At(x, y))
				diff = util.Weighted_hsv_distance(h1, s1, v1, h2, s2, v2, 1.0, 1.0, 1.0)
			} else {
				r, g, b, _ := refImg.At(x, y).RGBA()
				r1, g1, b1, _ := inputImg.At(x, y).RGBA()
				diff = util.ColorDifferenceRGB(color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}, color.RGBA{uint8(r1 >> 8), uint8(g1 >> 8), uint8(b1 >> 8), 255})
			}
			if diff < threshold {
				outputImg.Pix[offset+0] = 0
				outputImg.Pix[offset+1] = 0
				outputImg.Pix[offset+2] = 0
				outputImg.Pix[offset+3] = 255
			} else {
				outputImg.Pix[offset+0] = 255
				outputImg.Pix[offset+1] = 255
				outputImg.Pix[offset+2] = 255
				outputImg.Pix[offset+3] = 255
			}
		}
	}
}

func checkCompatibility(refImg image.Image, inputImg image.Image) error {
	if refImg.Bounds().Dx() != inputImg.Bounds().Dx() || refImg.Bounds().Dy() != inputImg.Bounds().Dy() {
		return fmt.Errorf("images are not compatible")
	}
	return nil
}