package main

import (
	"fmt"
	"math"
	"time"
)

type Screen struct {
	width       int
	height      int
	center_x    float64
	center_y    float64
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
	for i := 0; i < screen.width; i++ {
		for j := 0; j < screen.height; j++ {
			fmt.Printf("%c", screen.matrix.cells[i][j])
		}
		fmt.Println()
	}
}

func NewMatrix(width int, height int) Matrix {
	var m Matrix
	m.cells = make([][]rune, width)
	for i := 0; i < int(width); i++ {
		m.cells[i] = make([]rune, height)
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
	gradient := []rune{32, 46, 58, 33, 47, 114, 40, 108, 49, 90, 52, 72, 57, 87, 56, 36, 64}
	var camera Camera
	screen.center_x = float64(screen.width)/float64(2) - float64(1)
	screen.center_y = float64(screen.height)/float64(2) - float64(1)
	sphere := NewSphere(screen.center_x, screen.center_y, 0)
	for t := 0; t < times; t++ {
		for i := 0; i < screen.width; i++ {
			for j := 0; j < screen.height; j++ {
				camera.vector = Vector3D{-6, 0, 0}
				light := Vector3D{math.Sin(float64(t) * 0.01), -0.5, 1.0}
				intersection := Vector2D{0, 0}

				x := float64(i)/float64(screen.width)*float64(2) - float64(1)
				y := float64(j)/float64(screen.height)*float64(2) - float64(1)
				diff := math.Sin(sphere.z) * 16
				screen.matrix.cells[i][j] = gradient[1]

				if x*x+y*y < 0.5 {
					screen.matrix.cells[i][j] = gradient[int(diff)]
				}

				intersection = CollisionSphere(camera.vector.Sub(Vector3D{x: sphere.x, y: sphere.y, z: sphere.z}), light, sphere.r)
				fmt.Println(screen.center_x)
				if intersection.x > 0 {
					screen.matrix.cells[i][j] = gradient[int(15)]
				}
			}
		}
		screen.PrintScreen()
	}
}

func main() {
	var config Config
	config.width = 40
	config.height = 80
	fmt.Printf("\033c")
	screen := NewScreen(&config)
	screen.Animation(5)
}
