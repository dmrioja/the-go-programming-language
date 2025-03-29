package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

// exercise 1.12
// Modify the Lissajous server to read parameter values from the URL.
// For example, you might arrange it so that a URL like
// http://localhost:8000/?cycles=20 sets the number of cycles to 20
// instead of the default 5. Use strconv.Atoi function to convert the
// string parameter into an integer. You can see its documentation
// with go doc strconv.Atoi
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, r *http.Request) {
	const (
		whiteIndex = 0 // first color in var palette
		blackIndex = 1 // second color in var palette
	)
	var (
		palette = []color.Color{color.White, color.Black}
		cycles  = getQueryParam(r, "cycles", 5)        // number of complete x oscillator revolutions
		res     = getQueryParam(r, "res", 0.001)       // angular resolution
		size    = getQueryParam(r, "size", 100)        // image canvas covers [-size..+size]
		nframes = int(getQueryParam(r, "nframes", 64)) // number of animation frames
		delay   = int(getQueryParam(r, "delay", 8))    // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, int(2*size+1), int(2*size+1))
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(size+x*size+0.5), int(size+y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func getQueryParam(r *http.Request, queryParam string, defaultValue float64) float64 {
	qValue := r.URL.Query().Get(queryParam)

	if qValue != "" {
		v, _ := strconv.Atoi(qValue)
		return float64(v)
	}

	return defaultValue
}
