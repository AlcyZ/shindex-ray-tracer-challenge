package main

import (
	"strconv"
)

const Nl string = "\n"

type Canvas struct {
	Width  int
	Height int
	pixels [][]Color
}

func (c *Canvas) PixelAt(x int, y int) Color {
	return c.pixels[x][y]
}

func (c *Canvas) WritePixel(x int, y int, color Color) {
	c.pixels[x][y] = color
}

func NewCanvas(width int, height int) Canvas {
	pix := make([][]Color, width)
	black := Color{0, 0, 0}

	for i := 0; i < width; i++ {
		pix[i] = make([]Color, height)

		for j := 0; j < height; j++ {
			pix[i][j] = black
		}
	}

	return Canvas{
		Width:  width,
		Height: height,
		pixels: pix,
	}
}

func CanvasToPPM(c Canvas) string {
	return ppmHeaderLines(c)
}

func ppmHeaderLines(c Canvas) string {
	first := "P3"
	second := strconv.Itoa(c.Width) + " " + strconv.Itoa(c.Height)
	third := "255"

	return first + Nl + second + Nl + third
}
