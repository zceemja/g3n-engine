// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package math32 implements basic math functions which operate
// directly on float32 numbers without casting and contains
// types of common entities used in 3D Graphics such as vectors,
// matrices, quaternions and others.
package math32

import (
	m32 "github.com/chewxy/math32"
)

const Pi = m32.Pi
const degreeToRadiansFactor = m32.Pi / 180
const radianToDegreesFactor = 180.0 / m32.Pi

var Infinity = m32.Inf(1)

// DegToRad converts a number from degrees to radians
func DegToRad(degrees float32) float32 {

	return degrees * degreeToRadiansFactor
}

// RadToDeg converts a number from radians to degrees
func RadToDeg(radians float32) float32 {

	return radians * radianToDegreesFactor
}

// Clamp clamps x to the provided closed interval [a, b]
func Clamp(x, a, b float32) float32 {

	if x < a {
		return a
	}
	if x > b {
		return b
	}
	return x
}

// ClampInt clamps x to the provided closed interval [a, b]
func ClampInt(x, a, b int) int {

	if x < a {
		return a
	}
	if x > b {
		return b
	}
	return x
}

func Abs(v float32) float32 {
	return m32.Abs(v)
}

func Acos(v float32) float32 {
	return m32.Acos(v)
}

func Asin(v float32) float32 {
	return m32.Asin(v)
}

func Atan(v float32) float32 {
	return m32.Atan(v)
}

func Atan2(y, x float32) float32 {
	return m32.Atan2(y, x)
}

func Ceil(v float32) float32 {
	return m32.Ceil(v)
}

func Cos(v float32) float32 {
	return m32.Cos(v)
}

func Floor(v float32) float32 {
	return m32.Floor(v)
}

func Inf(sign int) float32 {
	return m32.Inf(sign)
}

func Round(v float32) float32 {
	return Floor(v + 0.5)
}

func IsNaN(v float32) bool {
	return m32.IsNaN(v)
}

func Sin(v float32) float32 {
	return m32.Sin(v)
}

func Sqrt(v float32) float32 {
	return m32.Sqrt(v)
}

func Log(v float32) float32 {
	return m32.Log(v)
}

func Exp(v float32) float32 {
	return m32.Exp(v)
}

func Max(a, b float32) float32 {
	return m32.Max(a, b)
}

func Min(a, b float32) float32 {
	return m32.Min(a, b)
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Mod(a, b float32) float32 {
	return m32.Mod(a, b)
}

func NaN() float32 {
	return m32.NaN()
}

func Pow(a, b float32) float32 {
	return m32.Pow(a, b)
}

func Tan(v float32) float32 {
	return m32.Tan(v)
}
