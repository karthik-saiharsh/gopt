package main

import "math"

/* In this file, we define a Vector.
 * A vector is primarily interpreted as an arrow starting from the origin,
 * and ending at point (x,y,z).
 * However, it can also be thought of as a single point (x,y,z) in space.
 */

type Vector struct {
	X, Y, Z float64
}

// NewVector creates a vector value.
func NewVector(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z}
}

// Add returns v + other.
func (v Vector) Add(other Vector) Vector {
	return Vector{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

// Sub returns v - other.
func (v Vector) Sub(other Vector) Vector {
	return Vector{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

// Scale returns v scaled by n.
func (v Vector) Scale(n float64) Vector {
	return Vector{X: v.X * n, Y: v.Y * n, Z: v.Z * n}
}

// Mul returns the component-wise product of two vectors.
func (v Vector) Mul(other Vector) Vector {
	return Vector{X: v.X * other.X, Y: v.Y * other.Y, Z: v.Z * other.Z}
}

// Length returns the Euclidean norm.
func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalized returns a unit-length copy of the vector.
func (v Vector) Normalized() Vector {
	length := v.Length()
	if length == 0 {
		return v
	}
	return v.Scale(1 / length)
}

// Normalize mutates the current vector to unit length.
func (v *Vector) Normalize() {
	*v = v.Normalized()
}

// Dot returns the dot product.
func (v Vector) Dot(other Vector) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross returns the cross product.
func (v Vector) Cross(other Vector) Vector {
	return Vector{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}
