package main

import (
	"math"
)

// Math Constants
const PI = math.Pi
const ONE_BY_PI = 1 / math.Pi

// The Scene
var Scene []Sphere = []Sphere{
	NewSphere(1e5, NewVector(1e5+1, 40.8, 81.6), NewVector(0, 0, 0), NewVector(0.75, 0.75, 0.75), DIFFUSE),   // Left
	NewSphere(1e5, NewVector(-1e5+99, 40.8, 81.6), NewVector(0, 0, 0), NewVector(0.25, 0.25, 0.75), DIFFUSE), // Right
	NewSphere(1e5, NewVector(50, 40.8, 1e5), NewVector(0, 0, 0), NewVector(0.75, 0.75, 0.75), DIFFUSE),       // Back
	NewSphere(1e5, NewVector(50, 40.8, -1e5+170), NewVector(0, 0, 0), NewVector(0, 0, 0), DIFFUSE),           // Front
	NewSphere(1e5, NewVector(50, 1e5, 81.6), NewVector(0, 0, 0), NewVector(0.75, 0.75, 0.75), DIFFUSE),       // Bottom
	NewSphere(1e5, NewVector(50, -1e5+81.6, 81.6), NewVector(0, 0, 0), NewVector(0.75, 0.75, 0.75), DIFFUSE), // Top
	NewSphere(16.5, NewVector(27, 16.5, 47), NewVector(0, 0, 0), NewVector(0.999, 0.999, 0.999), SPECULAR),   // Mirror
	NewSphere(16.5, NewVector(73, 16.5, 78), NewVector(0, 0, 0), NewVector(0.999, 0.999, 0.999), REFRACTIVE), // Glass
	NewSphere(1.5, NewVector(50, 81.6-16.5, 81.6), NewVector(400, 400, 400), NewVector(0, 0, 0), DIFFUSE),    // Light

}

var numSpheres int = len(Scene)

func main() {

}

// The output of the “radiance” function is a set of unbounded colors.
// This has to be converted to be between 0 and 255 for display purposes.

func clamp(x float64) float64 {
	if x < 0 {
		return 0
	} else if x > 1 {
		return 1
	} else {
		return x
	}
}

// This function applies a gamma correction of 2.2
func toInt(x float64) int {
	return int(math.Pow(clamp(x), 1/2.2)*255 + 0.5)
}

// This function calculates the intersection of a ray with the scene
func intersect(r *Ray, t *float64, id *int) bool {

	var inf float64 = 1e20

	for i := numSpheres; i != 0; i-- {
		var d float64 = Scene[i].Intersect(r)

		if d > 0 && d < *t {
			t = &d
			id = &i
		}
	}

	return *t < inf
}
