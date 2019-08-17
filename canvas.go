package main

import (
	"math"
	"strconv"
	"strings"
)

const (
	Nl string = "\n"
	Ws string = " "
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

	for x := 0; x < width; x++ {
		pix[x] = make([]Color, height)

		for y := 0; y < height; y++ {
			pix[x][y] = black
		}
	}

	return Canvas{
		Width:  width,
		Height: height,
		pixels: pix,
	}
}

func CanvasToPPM(c Canvas) string {
	hl := ppmHeaderLines(c)

	data := ""
	for y := 0; y < c.Height; y++ {
		line := ""
		for x := 0; x < c.Width; x++ {
			color := c.PixelAt(x, y)

			red := int(math.Ceil(color.Red() * 255))
			green := int(math.Ceil(color.Green() * 255))
			blue := int(math.Ceil(color.Blue() * 255))

			red = normRgb(red, 0, 255)
			green = normRgb(green, 0, 255)
			blue = normRgb(blue, 0, 255)

			r := strconv.Itoa(red)
			if len(line+r+Ws) > 70 {
				data = data + strings.TrimRight(line, Ws) + Nl + r + Ws
				line = ""
			} else {
				line = line + r + Ws
			}
			g := strconv.Itoa(green)
			if len(line+g+Ws) > 70 {
				data = data + strings.TrimRight(line, Ws) + Nl + g + Ws
				line = ""
			} else {
				line = line + g + Ws
			}
			b := strconv.Itoa(blue)
			if len(line+b+Ws) > 70 {
				data = data + strings.TrimRight(line, Ws) + Nl + b + Ws
				line = ""
			} else {
				line = line + b + Ws
			}
		}
		data = strings.TrimRight(data + line, Ws) + "\n"
	}

	return hl + Nl + data
}

func normRgb(val int, min int, max int) int {
	if val > 255 {
		val = 255
	}
	if val < 0 {
		val = 0
	}
	return val
}

func normLine(line string, color int) string {
	col := strconv.Itoa(color)
	newLine := line + col

	//fmt.Printf("%v\n\n", newLine)
	if len(newLine) > 70 {
		return line + Nl + col + Ws
	}

	return newLine + Ws
}

func ppmHeaderLines(c Canvas) string {
	first := "P3"
	second := strconv.Itoa(c.Width) + " " + strconv.Itoa(c.Height)
	third := "255"

	return first + Nl + second + Nl + third
}
