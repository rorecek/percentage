package percentage

import (
	"fmt"
	"testing"
)

func TestDifferenceBetween(t *testing.T) {
	var tests = []struct {
		a    float64
		b    float64
		want float64
	}{
		{100, 120, 20},
		{100, 80, -20},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v, %v", tt.a, tt.b)

		t.Run(name, func(t *testing.T) {
			got := DifferenceBetween(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsoluteDifferenceBetween(t *testing.T) {
	var tests = []struct {
		a    float64
		b    float64
		want float64
	}{
		{100, 120, 20},
		{100, 80, 20},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v, %v", tt.a, tt.b)

		t.Run(name, func(t *testing.T) {
			got := AbsoluteDifferenceBetween(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	var tests = []struct {
		new    float64
		origin float64
		want   float64
	}{
		{160, 80, 200},
		{60, 80, 75},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v, %v", tt.new, tt.origin)

		t.Run(name, func(t *testing.T) {
			got := Calculate(tt.new, tt.origin)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOf(t *testing.T) {
	var tests = []struct {
		percentage float64
		base       float64
		want       float64
	}{
		{20, 200, 40},
		{50, 200, 100},
		{-20, 200, -40},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v, %v", tt.percentage, tt.base)

		t.Run(name, func(t *testing.T) {
			got := Of(tt.percentage, tt.base)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleDifferenceBetween_positive() {
	fmt.Println(DifferenceBetween(200, 240))
	// Output: 20
}

func ExampleDifferenceBetween_negative() {
	fmt.Println(DifferenceBetween(200, 160))
	// Output: -20
}

func ExampleAbsoluteDifferenceBetween() {
	fmt.Println(AbsoluteDifferenceBetween(200, 160))
	// Output: 20
}

func ExampleCalculate() {
	fmt.Println(Calculate(160, 80))
	// Output: 200
}

func ExampleOf() {
	fmt.Println(Of(20, 200))
	// Output: 40
}
