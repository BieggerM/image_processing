package util

import (
	"image/color"
	"math"
)

/*
RgbToHsv is a function that converts an RGB color to an HSV color
It takes a color.Color as input and returns the hue, saturation and value of the color
*/
func RgbToHsv(color color.Color) (hue, saturation, value float64) {
	red, green, blue, _ := color.RGBA()
	// convert the RGB values to the range [0, 1]
	redF, greenF, blueF := float64(red)/65535.0, float64(green)/65535.0, float64(blue)/65535.0

	// find the maximum and minimum values of the RGB values
	max := math.Max(redF, math.Max(greenF, blueF))
	min := math.Min(redF, math.Min(greenF, blueF))
	delta := max - min

	if max == 0 {
		return 0, 0, 0 // Black color, hue is undefined
	}

	// value is max -> brightness
	value = max

	saturation = delta / max

	// hue calculation
	if redF == max {
		hue = (greenF - blueF) / delta
	} else if greenF == max {
		hue = 2 + (blueF-redF)/delta
	} else {
		hue = 4 + (redF-greenF)/delta
	}

	// convert hue to degrees
	hue *= 60
	if hue < 0 {
		hue += 360
	}

	return hue, saturation, value
}

func Weighted_hsv_distance(h1, s1, v1, h2, s2, v2, weightH, weightS, weightV float64) float64 {
	hDiff := h1 - h2
	sDiff := s1 - s2
	vDiff := v1 - v2
	return math.Sqrt(weightH*(hDiff*hDiff) + weightS*(sDiff*sDiff) + weightV*(vDiff*vDiff))
}
