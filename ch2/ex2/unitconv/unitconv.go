// Package unitconv performs unit conversions.
package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Foot float64
type Meter float64
type Pound float64
type Kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Foot) String() string       { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (p Pound) String() string      { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%gkg", k) }
