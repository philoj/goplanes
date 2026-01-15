package physics

import (
	geometry2 "github.com/philoj/goplanes/client/internal/geometry"

	"math"
)

type Mover interface {
	Location() geometry2.Vector
	Velocity() geometry2.Vector
	Heading() float64
	Move(delta float64)
	Rotate(dTheta float64)
	Turn(heading float64)
	Jump(location geometry2.Vector)
}

func NewMover(x, y, i, j, theta float64) Mover {
	return &movingObject{
		geometry2.Vector{I: x, J: y}, geometry2.Vector{I: i, J: j}, theta,
	}
}

type movingObject struct {
	location geometry2.Vector
	velocity geometry2.Vector
	heading  float64 // radians
}

func (p *movingObject) Location() geometry2.Vector {
	return p.location
}
func (p *movingObject) Velocity() geometry2.Vector {
	return p.velocity
}
func (p *movingObject) Heading() float64 {
	return p.heading
}
func (p *movingObject) Move(delta float64) {
	p.velocity.I, p.velocity.J = geometry2.RadialToXY(delta, p.heading)
	p.location.I += p.velocity.I
	p.location.J += p.velocity.J
}

func (p *movingObject) Rotate(dTheta float64) {
	p.heading += dTheta
	p.heading = math.Mod(p.heading, 2*math.Pi)
}
func (p *movingObject) Turn(heading float64) {
	p.heading = heading
}
func (p *movingObject) Jump(location geometry2.Vector) {
	p.location.I, p.location.J = location.I, location.J
}
