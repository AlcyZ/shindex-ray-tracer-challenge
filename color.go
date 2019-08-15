package main

type Color [3]float64

func (c *Color) Red() float64 {
	return c[0]
}

func (c *Color) Green() float64 {
	return c[1]
}

func (c *Color) Blue() float64 {
	return c[2]
}

func AddColors(a Color, b Color) Color {
	return Color{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

func SubtractColors(a Color, b Color) Color {
	return Color{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
}

func MultiplyColor(a Color, b float64) Color {
	return Color{
		a[0] * b,
		a[1] * b,
		a[2] * b,
	}
}

func MultiplyColors(a Color, b Color) Color {
	return Color{
		a[0] * b[0],
		a[1] * b[1],
		a[2] * b[2],
	}
}
