package main

import "math"

func (a Vector3D) RotateX(angle float64) Vector3D {
	b := a
	b.z = a.z*math.Cos(angle) - a.y*math.Sin(angle)
	b.y = a.z*math.Sin(angle) + a.y*math.Cos(angle)
	return b
}

func (a Vector3D) RotateY(angle float64) Vector3D {
	b := a
	b.x = a.x*math.Cos(angle) - a.z*math.Sin(angle)
	b.y = a.z*math.Sin(angle) + a.y*math.Cos(angle)
	return b
}

func (a Vector3D) RotateZ(angle float64) Vector3D {
	b := a
	b.x = a.x*math.Cos(angle) - a.y*math.Sin(angle)
	b.y = a.x*math.Sin(angle) + a.y*math.Cos(angle)
	return b
}
