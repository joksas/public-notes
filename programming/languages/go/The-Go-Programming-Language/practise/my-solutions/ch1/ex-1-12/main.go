// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
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

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.RequestURI()
	good_format := false
	if len(path) >= 9 {
		if path[:9] == "/?cycles=" {
			if len(path) > 9 {
				good_format = true
				cyclesStr := path[9:]
				cycles, err := strconv.ParseFloat(cyclesStr, 64)
				if err != nil {
					fmt.Fprintf(w, "The number of cycles should be... well, a number...")
				} else if cycles <= 0 {
					fmt.Fprintf(w, "The number of cycles should be larger than zero!")
				} else if cycles > 50 {
					fmt.Fprintf(w, "The number of cycles should not be larger than 50!")
				} else {
					lissajous(w, cycles)
				}
			}
		}
	}
	if good_format == false {
		fmt.Fprintf(w, "The URL should be \"http://localhost:8000/?cycles=num_cycles\" where num_cycles is the number of cycles.")
	}
}

func lissajous(out io.Writer, cycles float64) {
	const (
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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
