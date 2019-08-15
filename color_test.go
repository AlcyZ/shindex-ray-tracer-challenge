package main

import "testing"

func TestColorsAreRedGreenBlueTuples(t *testing.T) {
	c := Color{-0.5, 0.4, 1.7}

	if c.Red() != -0.5 {
		t.Errorf("Wrong red color: %v", c.Red())
	}
	if c.Green() != 0.4 {
		t.Errorf("Wrong red color: %v", c.Green())
	}
	if c.Blue() != 1.7 {
		t.Errorf("Wrong red color: %v", c.Blue())
	}
}

func TestAddingColors(t *testing.T) {
	a := Color{1, 2, 3}
	b := Color{2, 3, 4}

	expected := Color{3, 5, 7}
	actual := AddColors(a, b)

	if expected != actual {
		t.Errorf("expected did not match actual: %v - %v", expected, actual)
	}
}

func TestSubtractingColors(t *testing.T) {
	a := Color{2, 3, 4}
	b := Color{1, 2, 3}

	expected := Color{1, 1, 1}
	actual := SubtractColors(a, b)

	if expected != actual {
		t.Errorf("expected did not match actual: %v - %v", expected, actual)
	}
}

func TestMultiplyColorsByScalar(t *testing.T) {
	a := Color{1, 2, 3}

	expected := Color{3.5, 7, 10.5}
	actual := MultiplyColor(a, 3.5)

	if expected != actual {
		t.Errorf("expected did not match actual: %v - %v", expected, actual)
	}
}

func TestMultiplyColors(t *testing.T) {
	a := Color{1, 2, 3}
	b := Color{2, 3, 4}

	expected := Color{2, 6, 12}
	actual := MultiplyColors(a, b)

	if expected != actual {
		t.Errorf("expected did not match actual: %v - %v", expected, actual)
	}
}
