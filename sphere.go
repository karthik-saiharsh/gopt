package main

import "math"

type Material int

// Enum of surface materials
const (
	DIFFUSE    Material = iota
	SPECULAR   Material = iota
	REFRACTIVE Material = iota
)

type Sphere struct {
	radius                    float64
	position, emission, color Vector
	material                  Material
}

func NewSphere(r float64, p, e, c Vector, m Material) Sphere {
	return Sphere{
		radius:   r,
		position: p,
		emission: e,
		color:    c,
		material: m,
	}
}

func (s Sphere) Intersect(r Ray) float64 {
	op := s.position.Sub(r.origin)

	eps := 1e-4
	b := op.Dot(r.direction)
	det := b*b - op.Dot(op) + s.radius*s.radius

	if det < 0 {
		return 0
	}

	det = math.Sqrt(det)

	if b-det > eps {
		return b - det
	} else if b+det > eps {
		return b + det
	} else {
		return 0
	}
}
