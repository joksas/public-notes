// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320    // canvas size in pixels
	cells         = 100         // number of grid cells
	xyrange       = 30.0        // axis ranges (-xyrange..+xyrange)
	angle         = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	customHeight := paramToInt(r, "height", height)
	customWidth := paramToInt(r, "width", width)
	xyscale := float64(customWidth) / 2 / xyrange // pixels per x or y unit
	zscale := float64(customHeight) * 0.4         // pixels per z unit

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", customWidth, customHeight)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z := corner(i+1, j, customWidth, customHeight, xyscale, zscale)
			bx, by, _ := corner(i, j, customWidth, customHeight, xyscale, zscale)
			cx, cy, _ := corner(i, j+1, customWidth, customHeight, xyscale, zscale)
			dx, dy, _ := corner(i+1, j+1, customWidth, customHeight, xyscale, zscale)
			var color string
			if z > 0 {
				color = "#0072b2"
			} else {
				color = "#e69f00"
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func paramToInt(r *http.Request, paramName string, defaultInt int) int {
	customInt := defaultInt
	param := r.URL.Query()[paramName]
	if len(param) > 0 {
		i, err := strconv.Atoi(param[0])
		if err == nil {
			customInt = i
		}
	}

	return customInt
}

func corner(i, j int, customWidth, customHeight int, customxyscale, customzscale float64) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(customWidth)/2 + (x-y)*cos30*customxyscale
	sy := float64(customHeight)/2 + (x+y)*sin30*customxyscale - z*customzscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
