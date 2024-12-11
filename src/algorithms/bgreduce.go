package algorithms

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"time"
	"github.com/BieggerM/image_processing_golang/util"
)

func Background_reduction(reference string, input string, threshold float64, hsv bool) {
	start := time.Now()
	fmt.Println("-----Reading Reference Image-----")
	refImg, err := util.LoadImage(reference)
	if err != nil {
		fmt.Println("Failed to load reference image: ", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("[%s] Reference Image loaded \n", elapsed)

	fmt.Println("-----Reading Input Image-----")
	inputImg, err := util.LoadImage(input)
	if err != nil {
		fmt.Println("Failed to load input image: ", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("[%s] Input Image loaded \n", elapsed)

	fmt.Println("-----Checking compatibility of images-----")
	err = checkCompatibility(refImg, inputImg)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("[%s] Images are compatible \n", elapsed)

	fmt.Println("-----Background Reduction-----")
	outputImg := image.NewRGBA(refImg.Bounds())
		
	bounds := refImg.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if hsv {
				h1, s1, v1 := rgbToHsv(refImg.At(x, y))
				h2, s2, v2 := rgbToHsv(inputImg.At(x, y))
				diff := colorDifferenceHSV(h1, s1, v1, h2, s2, v2)
				if diff < threshold {
					outputImg.Set(x, y, color.RGBA{0, 0, 0, 255}) // Background
				} else {
					outputImg.Set(x, y, color.RGBA{255,255,255,255}) // Foreground
				}
			} else {
				r, g, b, _ := refImg.At(x, y).RGBA()
				r1, g1, b1, _ := inputImg.At(x, y).RGBA()
				diff := colorDifference(color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}, color.RGBA{uint8(r1 >> 8), uint8(g1 >> 8), uint8(b1 >> 8), 255})
				if diff < threshold {
					outputImg.Set(x, y, color.RGBA{0, 0, 0, 255}) // Background
				} else {
					outputImg.Set(x, y, color.RGBA{255,255,255,255}) // Foreground
				}
			}
		}
	}
	
	elapsed = time.Since(start)
	fmt.Printf("[%s] Background Reduction completed \n", elapsed)

	fmt.Println("-----Saving Output Image-----")
	err = util.SaveImage("output.jpg", outputImg)
	if err != nil {
		fmt.Println("Failed to save output image: ", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("[%s] Output Image saved in %s\n", elapsed, "output.jpg")
}


func checkCompatibility(refImg image.Image, inputImg image.Image) error {
	if refImg.Bounds().Dx() != inputImg.Bounds().Dx() || refImg.Bounds().Dy() != inputImg.Bounds().Dy() {
		return fmt.Errorf("images are not compatible")
	}
	return nil
}

func colorDifference(c1, c2 color.RGBA) float64 {
	rDiff := float64(c1.R) - float64(c2.R)
	gDiff := float64(c1.G) - float64(c2.G)
	bDiff := float64(c1.B) - float64(c2.B)
	return math.Sqrt((rDiff*rDiff + gDiff*gDiff + bDiff*bDiff) / 3.0)
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
