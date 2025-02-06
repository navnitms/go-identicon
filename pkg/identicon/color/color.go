package color

import (
	"image/color"
	"math"
)

type ColorGenerator interface {
	Generate(hash [16]byte) color.Color
}

type DefaultGenerator struct{}

func (g *DefaultGenerator) Generate(hash [16]byte) color.Color {
	hue := float64(hash[0]) / 255.0
	saturation := 0.5 + (float64(hash[1])/255.0)*0.5
	brightness := 0.4 + (float64(hash[2])/255.0)*0.4

	return hslToRGB(hue, saturation, brightness)
}

func hslToRGB(h, s, l float64) color.Color {
	var r, g, b float64

	if s == 0 {
		r, g, b = l, l, l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		r = hueToRGB(p, q, h+1.0/3.0)
		g = hueToRGB(p, q, h)
		b = hueToRGB(p, q, h-1.0/3.0)
	}

	return color.RGBA{
		R: uint8(math.Round(r * 255)),
		G: uint8(math.Round(g * 255)),
		B: uint8(math.Round(b * 255)),
		A: 255,
	}
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}
