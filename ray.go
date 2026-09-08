package main

/*
* This is a Ray.
* A Ray has a starting and ending point.
* Every Ray has an Origin Point (O) and a Direction (D)
* A point along a Ray can be denoted as P(t) = O + tD.
* O and D are Vectors, and t is a scalar
 */

type Ray struct {
	origin, direction Vector
}

func NewRay(o, d Vector) Ray {
	return Ray{origin: o, direction: d}
}
