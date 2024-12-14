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
	elapsed = time.Duration(time.Since(start).Seconds())
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
		numWorkers := 4 // Number of goroutines to use
		rowsPerWorker := (bounds.Max.Y - bounds.Min.Y) / numWorkers

		for w := 0; w < numWorkers; w++ {
			wg.Add(1)
			go func(worker int) {
				defer wg.Done()
				startY := bounds.Min.Y + worker*rowsPerWorker
				endY := startY + rowsPerWorker
				if worker == numWorkers-1 {
					endY = bounds.Max.Y
				}
				for y := startY; y < endY; y++ {
					for x := bounds.Min.X; x < bounds.Max.X; x++ {
						if hsv {
							h1, s1, v1 := rgbToHsv(refImg.At(x, y))
							h2, s2, v2 := rgbToHsv(inputImg.At(x, y))
							diff := colorDifferenceHSV(h1, s1, v1, h2, s2, v2)
							if diff < threshold {
								outputImg.Set(x, y, color.RGBA{0, 0, 0, 255}) // Background
							} else {
								outputImg.Set(x, y, color.RGBA{255, 255, 255, 255}) // Foreground
							}
						} else {
							r, g, b, _ := refImg.At(x, y).RGBA()
							r1, g1, b1, _ := inputImg.At(x, y).RGBA()
							diff := colorDifference(r, r1, g, g1, b, b1)
							if diff < threshold {
								outputImg.Set(x, y, color.RGBA{0, 0, 0, 255}) // Background
							} else {
								outputImg.Set(x, y, color.RGBA{255, 255, 255, 255}) // Foreground
							}
						}
					}
				}
			}(w)
		}
		wg.Wait()
	} else {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				if hsv {
					h1, s1, v1 := rgbToHsv(refImg.At(x, y))
					h2, s2, v2 := rgbToHsv(inputImg.At(x, y))
					diff := colorDifferenceHSV(h1, s1, v1, h2, s2, v2)
					if diff < threshold {
						outputImg.Set(x, y, color.RGBA{0, 0, 0, 255}) // Background
					} else {
						outputImg.Set(x, y, color.RGBA{255, 255, 255, 255}) // Foreground
					}
				} else {
					r, g, b, _ := refImg.At(x, y).RGBA()
					r1, g1, b1, _ := inputImg.At(x, y).RGBA()
					diff := colorDifference(r, r1, g, g1, b, b1)
					if diff < threshold {
						outputImg.Set(x, y, color.RGBA{0, 0, 0, 255}) // Background
					} else {
						outputImg.Set(x, y, color.RGBA{255, 255, 255, 255}) // Foreground
					}
				}
			}
		}
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

func checkCompatibility(refImg image.Image, inputImg image.Image) error {
	if refImg.Bounds().Dx() != inputImg.Bounds().Dx() || refImg.Bounds().Dy() != inputImg.Bounds().Dy() {
		return fmt.Errorf("images are not compatible")
	}
	return nil
}

func colorDifference(r1, r2, g1, g2, b1, b2 uint32) float64 {
	rDiff := r1 - r2
	gDiff := g1 - g2
	bDiff := b1 - b2
	return math.Sqrt(float64(rDiff*rDiff + gDiff*gDiff + bDiff*bDiff) / 3.0)
}

func colorDifferenceHSV(h1, s1, v1, h2, s2, v2 float64) float64 {
	hDiff := h1 - h2
	sDiff := s1 - s2
	vDiff := v1 - v2
	return math.Sqrt((hDiff*hDiff + sDiff*sDiff + vDiff*vDiff) / 3.0)
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
