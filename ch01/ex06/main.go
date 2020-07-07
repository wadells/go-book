// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Lissajous generates GIF animations of random Lissajous figures.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// sample return a palette of n colors evenly distributed across orig
func sample(orig []color.Color, n uint8) []color.Color {
	var np []color.Color
	for i := uint8(0); i < n; i++ {
		j := uint(float32(len(orig)) / float32(n) * float32(i))
		np = append(np, orig[j])
	}
	return np
}

func rainbowOnBlack() []color.Color {
	var rainbow, palette []color.Color
	var r, g, b, step, min, max uint8
	step = 1
	min = 0
	max = 255
	r = max                         // start at red
	for g = 0; g < max; g += step { // red to yellow
		rainbow = append(rainbow, color.RGBA{r, g, b, 0xff})
	}
	for ; r > min; r -= step { // yellow to green
		rainbow = append(rainbow, color.RGBA{r, g, b, 0xff})
	}
	for ; b < max; b += step { // yellow to cyan
		rainbow = append(rainbow, color.RGBA{r, g, b, 0xff})
	}
	for ; g > min; g -= step { // cyan to blue
		rainbow = append(rainbow, color.RGBA{r, g, b, 0xff})
	}
	for ; r < max; r += step { // blue to magenta
		rainbow = append(rainbow, color.RGBA{r, g, b, 0xff})
	}
	for ; b > min; b -= step { // magenta back to red
		rainbow = append(rainbow, color.RGBA{r, g, b, 0xff})
	}

	rainbow = sample(rainbow, 254)         // 256 capacity total, need 0 for black
	palette = append(palette, color.Black) // background
	palette = append(palette, rainbow...)
	return palette
}

var palette = rainbowOnBlack()

const (
	background = 0 // first color in palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(t*25)%uint8((len(palette)-1)) + 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
