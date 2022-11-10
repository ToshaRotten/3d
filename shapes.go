package main

import "math"

func Sphere(ro Vector3D, rd Vector3D, r float64) Vector2D {
	b := ro.Dot(rd)
	c := ro.Dot(ro) - r*r
	h := b*b - c
	if h < 0.0 {
		return Vector2D{
			x: -1.0,
			y: -1.0,
		}
	}
	h = math.Sqrt(h)
	return Vector2D{
		x: -b - h,
		y: -b + h,
	}
}
