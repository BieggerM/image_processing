package algorithms

import (
	"fmt"
	"image"
	"image/color"
	"math"
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
				h1, s1, v1 := rgbToHsv(refImg.At(x, y))
				h2, s2, v2 := rgbToHsv(inputImg.At(x, y))
				diff = weighted_hsv_distance(h1, s1, v1, h2, s2, v2, 1.0, 1.0, 1.0)
			} else {
				r, g, b, _ := refImg.At(x, y).RGBA()
				r1, g1, b1, _ := inputImg.At(x, y).RGBA()
				diff = colorDifferenceRGB(color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}, color.RGBA{uint8(r1 >> 8), uint8(g1 >> 8), uint8(b1 >> 8), 255})
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

func colorDifferenceRGB(c1, c2 color.RGBA) float64 {
	rDiff := float64(c1.R) - float64(c2.R)
	gDiff := float64(c1.G) - float64(c2.G)
	bDiff := float64(c1.B) - float64(c2.B)
	return math.Abs(rDiff + gDiff + bDiff) / 3.0
}

func rgbToHsv(c color.Color) (float64, float64, float64) {
	r, g, b, _ := c.RGBA()
	rf := float64(r) / 65535.0
	gf := float64(g) / 65535.0
	bf := float64(b) / 65535.0

	max := math.Max(rf, math.Max(gf, bf))
	min := math.Min(rf, math.Min(gf, bf))
	delta := max - min

	var h, s, v float64
	v = max

	if max != 0 {
		s = delta / max
	} else {
		s = 0
		h = -1
		return h, s, v
	}

	if rf == max {
		h = (gf - bf) / delta
	} else if gf == max {
		h = 2 + (bf-rf)/delta
	} else {
		h = 4 + (rf-gf)/delta
	}

	h *= 60
	if h < 0 {
		h += 360
	}

	return h, s, v
}

func weighted_hsv_distance(h1, s1, v1, h2, s2, v2, weightH, weightS, weightV float64) float64 {
	hDiff := h1 - h2
	sDiff := s1 - s2
	vDiff := v1 - v2
	return math.Sqrt(weightH*(hDiff*hDiff) + weightS*(sDiff*sDiff) + weightV*(vDiff*vDiff))
}