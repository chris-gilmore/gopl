// Package surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"fmt"
	"io"
	"math"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var svgColor = "grey"                         // stroke color
var svgWidth, svgHeight = 600, 320            // canvas size in pixels
var xyscale = float64(svgWidth) / 2 / xyrange // pixels per x or y unit
var zscale = float64(svgHeight) * 0.4         // pixels per z unit

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// SVG outputs an SVG rendering
func SVG(out io.Writer, color string, width, height int) {
	if color != "" {
		svgColor = color
	}
	if width > 0 {
		svgWidth = width
		xyscale = float64(svgWidth) / 2 / xyrange
	}
	if height > 0 {
		svgHeight = height
		zscale = float64(svgHeight) * 0.4
	}

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", svgColor, svgWidth, svgHeight)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(svgWidth)/2 + (x-y)*cos30*xyscale
	sy := float64(svgHeight)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
