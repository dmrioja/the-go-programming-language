package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
)

const (
	cycles  = 5     // number of complete x oscillator revolutions
	res     = 0.001 // angular resolution
	size    = 100   // image canvas covers [-size..+size]
	nFrames = 64    // number of animation frames
	delay   = 8     // delay between frames in 10ms units
)

func main() {
	exercise_1_5()
	exercise_1_6()
}

// exercise_1_5
// Change the Lissajous program`s color palette to green on black, for added
// authenticity. To create the web color #RRGGBB, use color.RGBA{0xRR, 0xGG, 0xBB, 0xFF},
// where each pair of hexadecimal digits represents the intensity of the red, green,
// or blue component of the pixel.
func exercise_1_5() {
	palette := []color.Color{color.Black, color.RGBA{R: 0, G: 255, B: 0, A: 255}}
	createLissajousGIF(palette, func(f float64) uint8 {
		return 1 // second color in var palette
	})
}

// exercise_1_6
// Modify the Lissajous program to produce images in multiple colors by adding
// more values to palette and then displaying them by changing the third argument
// of SetColorIndex in some interesting way.
func exercise_1_6() {
	palette := []color.Color{
		color.Black,
		// red colors
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
		color.RGBA{R: 200, G: 0, B: 0, A: 255},
		color.RGBA{R: 150, G: 0, B: 0, A: 255},
		color.RGBA{R: 100, G: 0, B: 0, A: 255},
		color.RGBA{R: 50, G: 0, B: 0, A: 255},
		// green colors
		color.RGBA{R: 0, G: 255, B: 0, A: 255},
		color.RGBA{R: 0, G: 200, B: 0, A: 255},
		color.RGBA{R: 0, G: 150, B: 0, A: 255},
		color.RGBA{R: 0, G: 100, B: 0, A: 255},
		color.RGBA{R: 0, G: 50, B: 0, A: 255},
		// blue colors
		color.RGBA{R: 0, G: 0, B: 255, A: 255},
		color.RGBA{R: 0, G: 0, B: 200, A: 255},
		color.RGBA{R: 0, G: 0, B: 150, A: 255},
		color.RGBA{R: 0, G: 0, B: 100, A: 255},
		color.RGBA{R: 0, G: 0, B: 50, A: 255},
	}
	createLissajousGIF(palette, func(f float64) uint8 {
		return uint8(int(f)%15) + 1
	})
}

func createLissajousGIF(palette color.Palette, iFn func(f float64) uint8) {

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0 // phase difference

	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), iFn(t))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(os.Stdout, &anim)
}
