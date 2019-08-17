package main

import (
	"fmt"
	"os"
)

type projectile struct {
	position Tuple
	velocity Tuple
}

type environment struct {
	gravity Tuple
	wind    Tuple
}

func tick(env environment, pro projectile, c Canvas) projectile {
	position := AddTuples(pro.position, pro.velocity)
	velocity := AddTuples(AddTuples(pro.velocity, env.gravity), env.wind)

	x := int(pro.position[0])
	y := c.Height - int(pro.position[1])
	col := Color{0.1, 0.9, 0}

	c.WritePixel(x, y, col)
	return projectile{position: position, velocity: velocity}
}

func main() {

	start := Point(0, 1, 0)
	vel := MultiplyTuples(NormalizeTuples(Vector(1, 1.8, 0)), 11.25)
	p := projectile{position: start, velocity: vel}

	gravity := Vector(0, -0.1, 0)
	wind := Vector(0.01, 0, 0)
	e := environment{gravity: gravity, wind: wind,}

	c := NewCanvas(900, 550)
	for p.position[1] >= 0 {
		p = tick(e, p, c)
	}

	fmt.Println("Creating ppm data...")
	ppm := CanvasToPPM(c)
	fmt.Println("Save data to file")
	//fmt.Printf(ppm)

	f, err := os.Create("test.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(ppm)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}
