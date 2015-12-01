package unitconv

// CelsiusToFahrenheit converts a Celsius temperature to Fahrenheit.
func CelsiusToFahrenheit(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FahrenheitToCelsius converts a Fahrenheit temperature to Celsius.
func FahrenheitToCelsius(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FootToMeter converts a Foot length to Meter.
func FootToMeter(f Foot) Meter { return Meter(f * 0.3048) }

// MeterToFoot converts a Meter length to Foot.
func MeterToFoot(m Meter) Foot { return Foot(m / 0.3048) }

// PoundToKilogram converts a Pound weight to Kilogram.
func PoundToKilogram(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KilogramToPound converts a Kilogram weight to Pound.
func KilogramToPound(k Kilogram) Pound { return Pound(k / 0.45359237) }
