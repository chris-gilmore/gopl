// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 2.0                 // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	minZ := 0.0
	maxZ := 0.0
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z, ok := ij2xyz(i, j)
			if !ok {
				continue
			}
			if z < minZ {
				minZ = z
			}
			if z > maxZ {
				maxZ = z
			}
		}
	}
	avgZ := (minZ + maxZ) / 2.0
	rangeZ := maxZ - minZ

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}

			_, _, z, ok := ij2xyz(i, j)
			f := z - avgZ
			var colorString string
			if f < 0.0 {
				// blue
				g := -f / (rangeZ / 2.0)
				colorString = fmt.Sprintf("#0000%02x", int(g*255))
			} else {
				// red
				g := f / (rangeZ / 2.0)
				colorString = fmt.Sprintf("#%02x0000", int(g*255))
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, colorString)
		}
	}
	fmt.Println("</svg>")
}

func ij2xyz(i, j int) (x, y, z float64, ok bool) {
	ok = true

	// Find point (x,y) at corner of cell (i,j).
	x = xyrange * (float64(i)/cells - 0.5)
	y = xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z = f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		ok = false
	}

	return
}

func corner(i, j int) (sx, sy float64, ok bool) {
	x, y, z, ok := ij2xyz(i, j)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale

	return
}

// f computes a monkey saddle
func f(x, y float64) float64 {
	return x*x*x - 3*x*y*y
}
