package main

import "math"

type Vertex struct {
	x float64
	y float64
	z float64
}

type Triangle struct {
	v1 Vector2D
	v2 Vector2D
	v3 Vector2D
}

type Sphere struct {
	x float64
	y float64
	z float64
	r float64
}

// New sphere
func NewSphere(x float64, y float64, z float64) *Sphere {
	return &Sphere{
		x: x,
		y: y,
		z: z,
		r: math.Sqrt(x*x + y*y + z*z),
	}
}

// New vertex
func NewVertex(x float64, y float64, z float64) *Vertex {
	return &Vertex{
		x: x,
		y: y,
		z: z,
	}
}

// New triangle
func NewTriangle(v1 Vector2D, v2 Vector2D, v3 Vector2D) *Triangle {
	return &Triangle{
		v1: v1,
		v2: v2,
		v3: v3,
	}
}

func CollisionSphere(v1 Vector3D, v2 Vector3D, r float64) Vector2D {
	b := v1.Dot(v2)
	c := v1.Dot(v1.Sub(Vector3D{r, r, r}.MultiplyByScalar(r)))
	h := b*b - c
	if h < 0.0 {
		return Vector2D{-1.0, -1.0}
	}
	h = math.Sqrt(h)
	return Vector2D{
		x: -b - h,
		y: -b + h,
	}
}
