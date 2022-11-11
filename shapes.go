package main

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
func NewSphere(x float64, y float64, z float64, r float64) *Sphere {
	return &Sphere{
		x: x,
		y: y,
		z: z,
		r: r,
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
