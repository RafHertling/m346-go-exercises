package main

import "fmt"

func main() {
	// Imperial to Metric conversions
	miles := 100.0
	fahrenheit := 80.0

	// Miles to Kilometers: 1 mile = 1.60934 kilometers
	kilometers := miles * 1.60934
	fmt.Printf("%.2f miles = %.2f kilometers\n", miles, kilometers)

	// Fahrenheit to Celsius: (°F - 32) × 5/9 = °C
	celsius := (fahrenheit - 32.0) * 5.0 / 9.0
	fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, celsius)

	// Metric to Imperial conversions
	kilometers2 := 50.0
	celsius2 := 25.0

	// Kilometers to Miles: 1 kilometer = 0.621371 miles
	miles2 := kilometers2 * 0.621371
	fmt.Printf("%.2f kilometers = %.2f miles\n", kilometers2, miles2)

	// Celsius to Fahrenheit: (°C × 9/5) + 32 = °F
	fahrenheit2 := (celsius2 * 9.0 / 5.0) + 32.0
	fmt.Printf("%.2f°C = %.2f°F\n", celsius2, fahrenheit2)
}
