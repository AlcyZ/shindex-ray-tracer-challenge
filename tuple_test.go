package shindex_ray_tracer_challenge

import (
	"math"
	"testing"
)

const (
	FailedTpl    = "Failed asserting that "
	EqualsTpl    = FailedTpl + "%v equals %v."
	NotEqualsTpl = FailedTpl + "%v not equals %v."
)

func TestATupleWithWEqualsOneIsAPoint(t *testing.T) {
	point := Tuple{1, 2, 3, 1}

	if !IsPoint(point) {
		t.Error(FailedTpl + "expected is point")
	}

	if IsVector(point) {
		t.Error(FailedTpl + "expected is not vector")
	}
}

func TestATupleWithWEqualsZeroIsAVector(t *testing.T) {
	vector := Tuple{1, 2, 3, 0}

	if !IsVector(vector) {
		t.Error(FailedTpl + "expected is vector")
	}

	if IsPoint(vector) {
		t.Error(FailedTpl + "expected is not")
	}
}

func TestFunctionPointCreatesTupleWithWEqualsOne(t *testing.T) {
	actual := Point(4, -4, 3)

	if actual[3] != 1 {
		t.Error("Point() function failed to created point (w=1)")
	}
}

func TestFunctionVectorCreatesTupleWithWEqualsZero(t *testing.T) {
	actual := Vector(4, -4, 3)

	if actual[3] != 0 {
		t.Error("Vector() function failed to created vector (w=0)")
	}
}

func TestFunctionCompareTuplesComparesTwoTuplesForEquality(t *testing.T) {
	tOne := Tuple{1, 1, 1, 1}
	tTwo := Tuple{1, 1, 0.999999, 1}

	tThree := Tuple{1, 1, 1, 1}
	tFor := Tuple{1, 1, 0.99990, 1}

	if !TupleEquals(tOne, tTwo) {
		t.Errorf(EqualsTpl, tOne, tTwo)
	}

	if TupleEquals(tThree, tFor) {
		t.Errorf(NotEqualsTpl, tThree, tFor)
	}
}

func TestAddingTwoTuples(t *testing.T) {
	a1 := Point(-2, 3, 1)
	a2 := Vector(3, -2, 5)
	expected := Tuple{1, 1, 6, 1}

	if a1.Add(a2) != expected {
		t.Error("Adding two tuples failed.")
	}
}

func TestSubtractingTwoPointsReturnsVector(t *testing.T) {
	a := Point(3, 2, 1)
	b := Point(5, 6, 7)
	expected := Vector(-2, -4, -6)
	actual := a.Subtract(b)

	if expected != actual {
		t.Errorf(EqualsTpl, expected, actual)
	}
}

func TestSubtractingAVectorFromAPoint(t *testing.T) {
	a := Point(3, 2, 1)
	b := Vector(5, 6, 7)
	expected := Point(-2, -4, -6)
	actual := a.Subtract(b)

	if expected != actual {
		t.Errorf(EqualsTpl, expected, actual)
	}
}

func TestSubtractingTwoVectors(t *testing.T) {
	a := Vector(3, 2, 1)
	b := Vector(5, 6, 7)
	expected := Vector(-2, -4, -6)
	actual := a.Subtract(b)

	if expected != actual {
		t.Errorf(EqualsTpl, expected, actual)
	}
}

func TestNegateVector(t *testing.T) {
	a := Vector(1, -2, 3)
	expected := Vector(-1, 2, -3)
	actual := a.Negate()

	if expected != actual {
		t.Errorf(EqualsTpl, expected, actual)
	}
}

func TestNegateTuple(t *testing.T) {
	a := Vector(1, -2, 3)
	expected := Vector(-1, 2, -3)
	actual := NegateTuple(a)

	if expected != actual {
		t.Errorf("%v negated not equals %v (actual: %v)", a, expected, actual)
	}
}

func TestMultiplyTuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	expected := Tuple{3.5, -7, 10.5, -14}
	actual := a.Multiply(3.5)

	if expected != actual {
		t.Errorf("%v multiplied by 3.5 not equals %v (actual: %v)", a, expected, actual)
	}
}

func TestDivideTuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	expected := Tuple{0.5, -1, 1.5, -2}
	actual := a.Divide(2)

	if expected != actual {
		t.Errorf("%v divided by 2 not equals %v (actual: %v)", a, expected, actual)
	}
}

type magnitudeDataSetHelper struct {
	v Tuple
	e float64
}

func TestComputingMagnitude(t *testing.T) {
	var expected []magnitudeDataSetHelper

	expected = append(expected, magnitudeDataSetHelper{Vector(1, 0, 0), 1})
	expected = append(expected, magnitudeDataSetHelper{Vector(0, 1, 0), 1})
	expected = append(expected, magnitudeDataSetHelper{Vector(0, 0, 1), 1})
	expected = append(expected, magnitudeDataSetHelper{Vector(1, 2, 3), math.Sqrt(14)})
	expected = append(expected, magnitudeDataSetHelper{Vector(-1, -2, -3), math.Sqrt(14)})

	for _, dataSet := range expected {
		actual := dataSet.v.Magnitude()
		expected := dataSet.e

		if expected != actual {
			t.Errorf("Failed to compute magnitude for vector %v\n(expected %v\nactual %v)", dataSet.v, expected, actual)
		}
	}
}

// unit vectors have magnitude equals one
func TestNormalizingTuplesReturnsUnitVector(t *testing.T) {
	var data []Tuple

	data = append(data, Vector(4, 0, 0))
	data = append(data, Vector(1, 2, 3))

	for _, vector := range data {
		normalized := vector.Normalize()
		actual := normalized.Magnitude()

		if actual != 1 {
			t.Errorf("Normalized vector have not magnitude of 1 (actual: %v).", actual)
		}
	}
}
