package util

import (
	"image/color"
	"math"
)

func RgbToHsv(c color.Color) (float64, float64, float64) {
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

func Weighted_hsv_distance(h1, s1, v1, h2, s2, v2, weightH, weightS, weightV float64) float64 {
	hDiff := h1 - h2
	sDiff := s1 - s2
	vDiff := v1 - v2
	return math.Sqrt(weightH*(hDiff*hDiff) + weightS*(sDiff*sDiff) + weightV*(vDiff*vDiff))
}

func ColorDifferenceRGB(c1, c2 color.RGBA) float64 {
	rDiff := float64(c1.R) - float64(c2.R)
	gDiff := float64(c1.G) - float64(c2.G)
	bDiff := float64(c1.B) - float64(c2.B)
	return math.Abs(rDiff + gDiff + bDiff) / 3.0
}

