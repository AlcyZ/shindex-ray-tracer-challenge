package main

import "math"

const EPSILON = 0.00001

type Tuple [4]float64

// arithmetic operations for tuple

func Add(a Tuple, b Tuple) Tuple {
	return Tuple{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
		a[3] + b[3],
	}
}

func Subtract(a Tuple, b Tuple) Tuple {
	return Tuple{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
		a[3] - b[3],
	}
}

func Multiply(t Tuple, x float64) Tuple {
	return Tuple{
		t[0] * x,
		t[1] * x,
		t[2] * x,
		t[3] * x,
	}
}

func Divide(t Tuple, x float64) Tuple {
	return Tuple{
		t[0] / x,
		t[1] / x,
		t[2] / x,
		t[3] / x,
	}
}

func Normalize(t Tuple) Tuple {
	return Tuple{
		t[0] / Magnitude(t),
		t[1] / Magnitude(t),
		t[2] / Magnitude(t),
		t[3] / Magnitude(t),
	}
}

func Magnitude(t Tuple) float64 {
	pow := math.Pow(t[0], 2) + math.Pow(t[1], 2) + math.Pow(t[2], 2) + math.Pow(t[3], 2)

	return math.Sqrt(pow)
}

func Dot(a Tuple, b Tuple) float64 {
	return a[0]*b[0] +
		a[1]*b[1] +
		a[2]*b[2] +
		a[3]*b[3]
}

func Cross(a Tuple, b Tuple) Tuple {
	x := a[1]*b[2] - a[2]*b[1]
	y := a[2]*b[0] - a[0]*b[2]
	z := a[0]*b[1] - a[1]*b[0]

	return Vector(x, y, z)
}

func Negate(t Tuple) Tuple {
	zero := Tuple{0, 0, 0, 0}

	return Subtract(zero, t)
}

// equality checks

func floatEquals(a float64, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

func Equals(tA Tuple, tB Tuple) bool {
	return floatEquals(tA[0], tB[0]) && floatEquals(tA[1], tB[1]) && floatEquals(tA[2], tB[2]) && tA[3] == tB[3]
}

// Sub-type comparisons

func IsPoint(t Tuple) bool {
	return t[3] == 1
}

func IsVector(t Tuple) bool {
	return t[3] == 0
}

// Factory methods for tuple.

func Point(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0}
}
