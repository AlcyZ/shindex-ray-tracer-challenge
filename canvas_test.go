package main

import (
	"strings"
	"testing"
)

func TestCreatingACanvas(t *testing.T) {
	c := NewCanvas(10, 20)

	if c.Width != 10 {
		t.Fail()
	}

	if c.Height != 20 {
		t.Fail()
	}
}

func TestNewCanvasHasEachPixelBlack(t *testing.T) {
	c := NewCanvas(3, 4)
	black := Color{0, 0, 0}

	for x := 0; x < 3; x++ {
		for y := 0; y < 4; y++ {
			if c.PixelAt(x, y) != black {
				t.Fail()
			}
		}
	}
}

func TestWritingPixelsToCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	red := Color{1, 0, 0}
	c.WritePixel(2, 3, red)

	if c.PixelAt(2, 3) != red {
		t.Fail()
	}
}

func TestConstructingPPMHeader(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := CanvasToPPM(c)

	expected := []string{"P3", "5 3", "255"}
	actual := strings.Split(ppm, "\n")

	if expected[0] != actual [0] {
		t.Errorf("First line is wrong, %v, %v", expected, actual)
	}
	if expected[1] != actual [1] {
		t.Errorf("Second line is wrong, %v, %v", expected, actual)
	}
	if expected[2] != actual [2] {
		t.Errorf("Third line is wrong, %v, %v", expected, actual)
	}
}

func TestConstructingThePPMPixelData(t *testing.T) {
	c := NewCanvas(5, 3)
	c1 := Color{1.5, 0, 0}
	c2 := Color{0, 0.5, 0}
	c3 := Color{-0.5, 0, 1}

	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)

	ppm := CanvasToPPM(c)

	l1 := "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0"
	l2 := "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0"
	l3 := "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255"

	expected := []string{l1, l2, l3}
	actual := strings.Split(ppm, "\n")

	if expected[0] != actual [3] {
		t.Errorf("Fourth line is wrong, %v, %v", expected[0], actual[3])
	}
	if expected[1] != actual [4] {
		t.Errorf("Fifth line is wrong, %v, %v", expected[1], actual[4])
	}
	if expected[2] != actual [5] {
		t.Errorf("Sixth line is wrong, %v, %v", expected[2], actual[5])
	}
}
