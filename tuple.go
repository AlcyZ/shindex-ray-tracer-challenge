package shindex_ray_tracer_challenge

import "math"

const EPSILON = 0.00001

type Tuple [4]float64

// arithmetic operations for tuple

func (t *Tuple) Add(other Tuple) Tuple {
	return Tuple{
		t[0] + other[0],
		t[1] + other[1],
		t[2] + other[2],
		t[3] + other[3],
	}
}

func (t *Tuple) Subtract(other Tuple) Tuple {
	return Tuple{
		t[0] - other[0],
		t[1] - other[1],
		t[2] - other[2],
		t[3] - other[3],
	}
}

func (t *Tuple) Multiply(x float64) Tuple {
	return Tuple{
		t[0] * x,
		t[1] * x,
		t[2] * x,
		t[3] * x,
	}
}

func (t *Tuple) Divide(x float64) Tuple {
	return Tuple{
		t[0] / x,
		t[1] / x,
		t[2] / x,
		t[3] / x,
	}
}

func (t *Tuple) Normalize () Tuple {
	return Tuple{
		t[0] / t.Magnitude(),
		t[1] / t.Magnitude(),
		t[2] / t.Magnitude(),
		t[3] / t.Magnitude(),
	}
}

func (t *Tuple) Magnitude() float64 {
	pow := math.Pow(t[0], 2) + math.Pow(t[1], 2) + math.Pow(t[2], 2) + math.Pow(t[3], 2)

	return math.Sqrt(pow)
}

func (t *Tuple) Negate() Tuple {
	zero := Tuple{0, 0, 0, 0}

	return zero.Subtract(*t)
}

func NegateTuple(t Tuple) Tuple {
	zero := Tuple{0, 0, 0, 0}

	return zero.Subtract(t)
}

// equality checks

func floatEquals(a float64, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

func TupleEquals(tA Tuple, tB Tuple) bool {
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
