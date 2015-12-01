package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/chris-gilmore/gopl/ch2/ex2/unitconv"
)

func main() {
	if len(os.Args[1:]) > 0 {
		for _, arg := range os.Args[1:] {
			v, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				continue
			}

			convert(v)
		}
	} else {
		for {
			var v float64
			_, err := fmt.Scan(&v)
			if err != nil {
				if err != io.EOF {
					continue
				}
				break
			}

			convert(v)
		}
	}
}

func convert(v float64) {
	fahrenheit := unitconv.Fahrenheit(v)
	celsius := unitconv.Celsius(v)
	foot := unitconv.Foot(v)
	meter := unitconv.Meter(v)
	pound := unitconv.Pound(v)
	kilogram := unitconv.Kilogram(v)

	fmt.Printf("%s = %s, %s = %s; %s = %s, %s = %s; %s = %s, %s = %s\n",
		fahrenheit, unitconv.FahrenheitToCelsius(fahrenheit), celsius, unitconv.CelsiusToFahrenheit(celsius),
		foot, unitconv.FootToMeter(foot), meter, unitconv.MeterToFoot(meter),
		pound, unitconv.PoundToKilogram(pound), kilogram, unitconv.KilogramToPound(kilogram))
}
