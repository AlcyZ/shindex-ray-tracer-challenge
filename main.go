package main

import "fmt"

type projectile struct {
	position Tuple
	velocity Tuple
}

type environment struct {
	gravity Tuple
	wind    Tuple
}

func tick(env environment, pro projectile) projectile {
	position := AddTuples(pro.position, pro.velocity)
	velocity := AddTuples(AddTuples(pro.velocity, env.gravity), env.wind)

	return projectile{position: position, velocity: velocity}
}

func main() {
	pro := projectile{
		position: Point(0, 1, 0),
		velocity: NormalizeTuples(Vector(1, 1, 1)),
	}

	env := environment{
		gravity: Vector(0, -1, 0),
		wind:    Vector(-0.01, 0, 0),
	}

	for pro.position[1] >= 0 {
		x := pro.position[0]
		y := pro.position[1]
		z := pro.position[2]

		fmt.Printf("\nx: %v\ny: %v\nz: %v\n", x, y, z)
		pro = tick(env, pro)
	}
}
