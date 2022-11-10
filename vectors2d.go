package main

import "math"

// Type of 2d-vector
type Vector2D struct {
	x float64
	y float64
}

// Addintion of two 2d-Vectors
func (a Vector2D) Add(b Vector2D) Vector2D {
	a.x += b.x
	a.y += b.y
	return a
}

// Subtraction of two 2d-Vectors
func (a Vector2D) Sub(b Vector2D) Vector2D {
	return Vector2D{
		x: a.x - b.x,
		y: a.y - b.y,
	}
}

// Multiplying by a scalar
func (a Vector2D) MultiplyByScalar(s float64) Vector2D {
	return Vector2D{
		x: a.x * s,
		y: a.y * s,
	}
}

// These case give us a way to determine if two rays are orthogonal or codirectional
func (a Vector2D) Dot(b Vector2D) float64 {
	return a.x*b.x + a.y*b.y
}

// Algebraic vector length
func (a Vector2D) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

// Normalising a Vector
func (a Vector2D) Normalize() Vector2D {
	return a.MultiplyByScalar(1. / a.Length())
}
