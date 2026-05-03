package main

import "math"

/* In this file, we define a Vector.
* A vector is primarily interpreted as an arrow starting from the origin,
* and ending at point (x,y,z).
* However, it can also be thought of as a single point (x,y,x) in space.
 */

type Vector struct {
	X, Y, Z float64
}

// New Vector Object
func NewVector(x, y, z float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

// Add Operation
func (this *Vector) Add(v Vector) *Vector {
	return &Vector{
		X: this.X + v.X,
		Y: this.Y + v.Y,
		Z: this.Z + v.Z,
	}
}

// Subtract Operation
func (this *Vector) Sub(v Vector) *Vector {
	return &Vector{
		X: this.X - v.X,
		Y: this.Y - v.Y,
		Z: this.Z - v.Z,
	}
}

// Scale Operation
func (this *Vector) Scale(n float64) *Vector {
	return &Vector{
		X: this.X * n,
		Y: this.Y * n,
		Z: this.Z * n,
	}
}

// Direct Multiplication Operation
func (this *Vector) Mul(v Vector) *Vector {
	return &Vector{
		X: this.X * v.X,
		Y: this.Y * v.Y,
		Z: this.Z * v.Z,
	}
}

// Normalize the current vector
func (this *Vector) Normalize() {
	var scaleVal float64 = 1 / math.Sqrt(this.X*this.X+this.Y*this.Y+this.Z*this.Z)
	this = this.Scale(scaleVal)
}

// Dot Product
func (this *Vector) Dot(v Vector) float64 {
	return this.X*v.X + this.Y*v.Y + this.Z*v.Z
}

// Cross Product
func (this *Vector) Cross(v Vector) *Vector {
	return &Vector{
		X: this.Y*v.Z - this.Z*v.Y,
		Y: this.Z*v.X - this.X*v.Z,
		Z: this.X*v.Y - this.Y*v.X,
	}
}
