package main

import "math"

// Type of 3d-vector
type Vector3D struct {
	x float64
	y float64
	z float64
}

// Addintion of two 3d-Vectors
func (a Vector3D) Add(b Vector3D) Vector3D {
	a.x += b.x
	a.y += b.y
	a.z += b.z
	return a
}

// Subtraction of two 3d-Vectors
func (a Vector3D) Sub(b Vector3D) Vector3D {
	return Vector3D{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
	}
}

// Multiplying by a scalar
func (a Vector3D) MultiplyByScalar(s float64) Vector3D {
	return Vector3D{
		x: a.x * s,
		y: a.y * s,
		z: a.z * s,
	}
}

// These two cases give us a way to determine if two rays are orthogonal or codirectional
func (a Vector3D) Dot(b Vector3D) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

// Algebraic vector length
func (a Vector3D) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

// Normalising a Vector
func (a Vector3D) Normalize() Vector3D {
	return a.MultiplyByScalar(1. / a.Length())
}
