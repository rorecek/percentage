// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
A clean interface to calculate percentages in Go: changes between values, percentage increases and partial values.
Think of it as a very lightweight math library for the basic stuff. Aimed to make your code more readable and easier to understand.

Usage

What's the percentage increase from 100 to 120?
 percentage.DifferenceBetween(100, 120);         // 20

What's the absolute percentage change from 100 to 80?
 percentage.AbsoluteDifferenceBetween(100, 80);  // 20, not -20

How much is 120 compared to 100?
 percentage.Calculate(120, 100);                 // 120%

How much is 50 compared to 100?
 percentage.Calculate(50, 100);                  // 50%

What is 20% of 200?
 percentage.Of(20, 200);                         // 40
*/
package percentage

import "math"

// What is the percentage increase or decrease from a to b
func DifferenceBetween(a, b float64) float64 {
	return (b - a) / a * 100
}

// What is the absolute percentage increase or decrease from a to b
func AbsoluteDifferenceBetween(a, b float64) float64 {
	return math.Abs(DifferenceBetween(a, b))
}

// How much is the new value of the origin value in percentages
func Calculate(new, origin float64) float64 {
	return new / origin * 100
}

// Get a percentage from the base number
func Of(percentage, base float64) float64 {
	return percentage * base / 100
}
