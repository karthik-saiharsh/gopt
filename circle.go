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

func NewSphere(r float64, p, e, c Vector, m Material) *Sphere {
	return &Sphere{
		radius:   r,
		position: p,
		emission: e,
		color:    c,
		material: m,
	}
}

func (this *Sphere) Intersect(r *Ray) float64 {
	var op Vector = *this.position.Sub(r.origin)

	var eps float64 = 1e-4
	var b float64 = op.Dot(r.direction)
	var det float64 = b*b - op.Dot(op) + this.radius*this.radius

	if det < 0 {
		return 0
	} else {
		det = math.Sqrt(det)
	}

	if b-det > eps {
		return b - det
	} else if b+det > eps {
		return b + det
	} else {
		return 0
	}
}
