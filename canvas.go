package main

import (
	"strconv"
	"strings"
)

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
	lines := make([]string, 3 + c.Height)

	lines[0] = "P3"
	lines[1] = strconv.Itoa(c.Width) + " " + strconv.Itoa(c.Height)
	lines[2] = "255"

	//for i := 0; i < c.Height; i++ {
	//	line := make([]string, c.Width)
	//}

	return strings.Join(lines, "\n")
}
