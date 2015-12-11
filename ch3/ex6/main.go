// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			var c color.RGBA
			x := float64(px)/width*(xmax-xmin) + xmin
			c = adjustColor(c, mandelbrot(complex(x, y)), 4)
			x1 := float64(px+1)/width*(xmax-xmin) + xmin
			c = adjustColor(c, mandelbrot(complex(x1, y)), 4)
			y1 := float64(py+1)/height*(ymax-ymin) + ymin
			c = adjustColor(c, mandelbrot(complex(x1, y1)), 4)
			c = adjustColor(c, mandelbrot(complex(x, y1)), 4)

			// Image point (px, py) represents complex value z.
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func adjustColor(ca color.RGBA, cb color.Color, n uint8) color.RGBA {
	r, g, b, a := ca.RGBA()
	retR := uint8(r)
	retG := uint8(g)
	retB := uint8(b)
	retA := uint8(a)

	r, g, b, a = cb.RGBA()
	retR += uint8(r) / n
	retG += uint8(g) / n
	retB += uint8(b) / n
	retA += uint8(a) / n

	return color.RGBA{retR, retG, retB, retA}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			return palette.Plan9[255-contrast*n]
		}
	}
	return color.Black
}
