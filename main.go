package main

import (
	"fmt"
	"math"
	"time"
)

type Screen struct {
	width       int
	height      int
	aspect      float64
	pixelAspect float64
	matrix      Matrix
}

type Matrix struct {
	cells [][]rune
}

type Config struct {
	width  int
	height int
}

func (screen *Screen) PrintScreen() {
	time.Sleep(time.Second / 18)
	for i := 0; i < screen.height; i++ {
		for j := 0; j < screen.width; j++ {
			fmt.Printf("%c", screen.matrix.cells[i][j])
		}
		fmt.Println()
	}
}

func NewMatrix(width int, height int) Matrix {
	var m Matrix
	m.cells = make([][]rune, height)
	for i := 0; i < int(height); i++ {
		m.cells[i] = make([]rune, width)
	}
	return m
}

func NewAspect(width int, height int) float64 {
	return float64(width) / float64(height)
}

func NewPixelAspect(width int, height int) float64 {
	return float64(11) / float64(24)
}

func NewScreen(config *Config) *Screen {
	return &Screen{
		width:       config.width,
		height:      config.height,
		aspect:      NewAspect(config.width, config.height),
		pixelAspect: NewPixelAspect(config.width, config.height),
		matrix:      NewMatrix(config.width, config.height),
	}
}

func (screen *Screen) Animation(times int) {
	var color float64
	uv := Vector2D{0, 0}
	// Camera position
	ro := Vector3D{1, 1, 1}
	rd := Vector3D{1, 1, 1}
	light := Vector3D{1, 1, 1}
	itPoint := Vector3D{0, 0, 0}

	n := Vector3D{1, 1, 1}
	var diff float64

	gradient := []rune{32, 46, 58, 33, 47, 114, 40, 108, 49, 90, 52, 72, 57, 87, 56, 36, 64}
	for t := 0; t < times; t++ {
		for i := 0; i < screen.height; i++ {
			for j := 0; j < screen.width; j++ {
				uv = Vector2D{float64(i) / float64(screen.height), float64(j) / float64(screen.width)}
				uv = uv.MultiplyByScalar(2)
				uv = uv.Sub(Vector2D{1.0, 1.0})
				uv.y *= screen.aspect * screen.pixelAspect
				ro = Vector3D{-8, 0, 0}
				rd = Vector3D{2, uv.x, uv.y}
				rd = rd.Normalize()

				light = Vector3D{math.Sin((float64(t)) * 0.001), math.Cos((float64(t) * 0.001)), -1.0}
				light = light.Normalize()

				var intersection Vector2D
				intersection = Sphere(ro, rd, 3)

				if intersection.x > 0 {
					itPoint = itPoint.Add(ro).Add(rd).MultiplyByScalar(intersection.x)
					n = itPoint.Normalize()
					diff *= n.Dot(light)
					color = diff * 15
					fmt.Print(n.Dot(light))
				} else {
					color = 1
				}

				color = clamp(color, 0, float64(len(gradient)))
				screen.matrix.cells[i][j] = gradient[int(color)]

			}
		}
		screen.PrintScreen()
	}
}

func main() {
	var config Config
	config.width = 120
	config.height = 30
	fmt.Printf("\033c")
	screen := NewScreen(&config)
	screen.Animation(500)
}
