// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

// http://localhost:8000/?x=-1.474&y=0&zoom=17

// http://localhost:8000/?x=-1.47468812&y=0&zoom=31&iter=3000

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			var cx, cy float64
			var zoom int
			var iter int = 256

			if v, err := strconv.ParseFloat(r.FormValue("x"), 64); err == nil {
				cx = v
			}
			if v, err := strconv.ParseFloat(r.FormValue("y"), 64); err == nil {
				cy = v
			}
			if v, err := strconv.Atoi(r.FormValue("zoom")); err == nil {
				zoom = v
			}
			if v, err := strconv.Atoi(r.FormValue("iter")); err == nil {
				iter = v
			}
			fractal2png(w, cx, cy, zoom, iter)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	fractal2png(os.Stdout, 0, 0, 0, 256)
}

func fractal2png(out io.Writer, cx, cy float64, zoom int, iter int) {
	const (
		width, height = 1500, 600
	)

	var xmin, ymin, xmax, ymax float64 = -3, -1.2, +3, +1.2

	scale := math.Pow(2, float64(-zoom))

	xmin *= scale
	ymin *= scale
	xmax *= scale
	ymax *= scale

	xmin += cx
	ymin += cy
	xmax += cx
	ymax += cy

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z, iter))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128, iterations int) color.Color {
	const contrast = 1

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			return palette.Plan9[255-(contrast*n)%256]
		}
	}
	return color.Black
}
