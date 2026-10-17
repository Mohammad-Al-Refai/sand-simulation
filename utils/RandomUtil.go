package utils

import (
	"image/color"
	"math/rand/v2"
)

func RandomBool() bool {
	return rand.Int64N(2) == 1
}

func RandomColor(colors []color.RGBA) color.RGBA {
	return colors[rand.IntN(len(colors))]
}
