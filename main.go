package main

import (
	"fmt"
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
	for t := 0; t < times; t++ {
		for i := 0; i < screen.width; i++ {
			for j := 0; j < screen.height; j++ {
				x := float64(i) / float64(screen.width) * float64(2)
				y := float64(j) / float64(screen.height) * float64(2)

				screen.matrix.cells[i][j] = gradient[1]
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
	screen.Animation(1)
}
