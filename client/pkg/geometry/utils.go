package geometry

import (
	"fmt"
	"math"
)

func RadialToXY(radius, theta float64) (x, y float64) {
	return radius * math.Cos(theta), radius * math.Sin(theta)
}

func BisectRectangle(p1, p2, rectMin, rectMax Vector) Vector {
	p := Vector{
		I: p2.I,
		J: p2.J,
	}
	if rectMin.I > rectMax.I || rectMin.J > rectMax.J {
		panic(fmt.Errorf("invalid values for min and max %v, %v", rectMin, rectMax))
	}
	if p1.I > rectMin.I && p1.I < rectMax.I && p1.J > rectMin.J && p1.J < rectMax.J {
		if p2.I > rectMax.I {
			p.I = rectMax.I
		} else if p2.I < rectMin.I {
			p.I = rectMin.I
		}
		if p2.J > rectMax.J {
			p.J = rectMax.J
		} else if p2.J < rectMin.J {
			p.J = rectMin.J
		}
		if p.I == p2.I && p.J == p2.J {
			panic("invalid value for p2")
		}
		return p
	}
	panic("invalid value for p1")
}

func AxialDistance(p1, p2 Vector) Vector {
	return Vector{I: p2.I - p1.I, J: p2.J - p1.J}
}

func Theta(v Vector) float64 {
	tan := v.J / v.I
	if tan == 0 {
		if v.I > 0 {
			return 0
		} else {
			return math.Pi
		}
	} else if tan < 0 {
		theta := math.Atan(-tan)
		if v.J > 0 {
			return math.Pi - theta
		} else {
			return -theta
		}
	} else {
		theta := math.Atan(tan)
		if v.J < 0 {
			return math.Pi + theta
		} else {
			return theta
		}
	}
}

func Degrees(rad float64) float64 {
	return rad * 180 / math.Pi
}
