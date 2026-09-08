package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

// Math Constants
const PI = math.Pi
const ONE_BY_PI = 1 / math.Pi

const (
	width  = 1024
	height = 768
)

var samps = flag.Int("samps", 50, "samples per subpixel")

// The Scene
var Scene = []Sphere{
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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintf(writer, "P3\n%d %d\n%d\n", width, height, 255)

	cam := Ray{
		origin:    NewVector(50, 52, 295.6),
		direction: NewVector(0, -0.042612, -1).Normalized(),
	}
	cx := NewVector(float64(width)*0.5135/float64(height), 0, 0)
	cy := cx.Cross(cam.direction).Normalized().Scale(0.5135)
	cx = cx.Normalized().Scale(0.5135)

	for y := height - 1; y >= 0; y-- {
		fmt.Fprintf(os.Stderr, "\rRendering %5.2f%%", 100*float64(height-1-y)/float64(height-1))
		for x := 0; x < width; x++ {
			pixel := NewVector(0, 0, 0)
			for sy := 0; sy < 2; sy++ {
				for sx := 0; sx < 2; sx++ {
					subpixel := NewVector(0, 0, 0)
					for s := 0; s < *samps; s++ {
						dx := tentFilter()
						dy := tentFilter()
						d := cx.Scale(((float64(sx)+0.5+dx)/2+float64(x))/float64(width) - 0.5).
							Add(cy.Scale(((float64(sy)+0.5+dy)/2+float64(y))/float64(height) - 0.5)).
							Add(cam.direction)
						subpixel = subpixel.Add(radiance(NewRay(cam.origin.Add(d.Scale(140)), d.Normalized()), 0).Scale(1 / float64(*samps)))
					}
					pixel = pixel.Add(subpixel.Scale(0.25))
				}
			}
			fmt.Fprintf(writer, "%d %d %d ", toInt(pixel.X), toInt(pixel.Y), toInt(pixel.Z))
		}
		fmt.Fprintln(writer)
	}
	fmt.Fprintln(os.Stderr)
}

// The output of the radiance function is a set of unbounded colors.
// This has to be converted to be between 0 and 255 for display purposes.
func clamp(x float64) float64 {
	if x < 0 {
		return 0
	}
	if x > 1 {
		return 1
	}
	return x
}

// This function applies a gamma correction of 2.2.
func toInt(x float64) int {
	return int(math.Pow(clamp(x), 1/2.2)*255 + 0.5)
}

func tentFilter() float64 {
	r := 2 * rand.Float64()
	if r < 1 {
		return math.Sqrt(r) - 1
	}
	return 1 - math.Sqrt(2-r)
}

// This function calculates the intersection of a ray with the scene.
func intersect(r Ray) (bool, float64, int) {
	closest := 1e20
	hitID := -1

	for i, sphere := range Scene {
		if d := sphere.Intersect(r); d > 0 && d < closest {
			closest = d
			hitID = i
		}
	}

	return hitID != -1, closest, hitID
}

func radiance(r Ray, depth int) Vector {
	hit, t, id := intersect(r)
	if !hit {
		return NewVector(0, 0, 0)
	}

	obj := Scene[id]
	x := r.origin.Add(r.direction.Scale(t))
	n := x.Sub(obj.position).Normalized()
	nl := n
	if n.Dot(r.direction) > 0 {
		nl = n.Scale(-1)
	}
	f := obj.color
	depth++

	if depth > 5 {
		p := math.Max(f.X, math.Max(f.Y, f.Z))
		if rand.Float64() < p {
			f = f.Scale(1 / p)
		} else {
			return obj.emission
		}
	}

	switch obj.material {
	case DIFFUSE:
		r1 := 2 * PI * rand.Float64()
		r2 := rand.Float64()
		r2s := math.Sqrt(r2)

		w := nl
		u := NewVector(0, 1, 0).Cross(w)
		if math.Abs(w.X) <= 0.1 {
			u = NewVector(1, 0, 0).Cross(w)
		}
		u = u.Normalized()
		v := w.Cross(u)
		d := u.Scale(math.Cos(r1) * r2s).
			Add(v.Scale(math.Sin(r1) * r2s)).
			Add(w.Scale(math.Sqrt(1 - r2))).Normalized()

		return obj.emission.Add(f.Mul(radiance(NewRay(x, d), depth)))
	case SPECULAR:
		d := r.direction.Sub(n.Scale(2 * n.Dot(r.direction))).Normalized()
		return obj.emission.Add(f.Mul(radiance(NewRay(x, d), depth)))
	case REFRACTIVE:
		reflDir := r.direction.Sub(n.Scale(2 * n.Dot(r.direction))).Normalized()
		reflRay := NewRay(x, reflDir)

		into := n.Dot(nl) > 0
		nc := 1.0
		nT := 1.5
		nnt := nc / nT
		if !into {
			nnt = nT / nc
		}
		ddn := r.direction.Dot(nl)
		cos2t := 1 - nnt*nnt*(1-ddn*ddn)
		if cos2t < 0 {
			return obj.emission.Add(f.Mul(radiance(reflRay, depth)))
		}

		tdir := r.direction.Scale(nnt)
		term := nnt*ddn + math.Sqrt(cos2t)
		if into {
			tdir = tdir.Sub(n.Scale(term))
		} else {
			tdir = tdir.Add(n.Scale(term))
		}
		tdir = tdir.Normalized()
		transRay := NewRay(x, tdir)

		a := nT - nc
		b := nT + nc
		R0 := (a * a) / (b * b)
		c := 1 - func() float64 {
			if into {
				return -ddn
			}
			return tdir.Dot(n)
		}()
		Re := R0 + (1-R0)*math.Pow(c, 5)
		Tr := 1 - Re

		if depth > 2 {
			P := 0.25 + 0.5*Re
			if rand.Float64() < P {
				return obj.emission.Add(f.Mul(radiance(reflRay, depth).Scale(Re / P)))
			}
			return obj.emission.Add(f.Mul(radiance(transRay, depth).Scale(Tr / (1 - P))))
		}

		return obj.emission.Add(
			f.Mul(radiance(reflRay, depth).Scale(Re).Add(radiance(transRay, depth).Scale(Tr))),
		)
	default:
		return obj.emission
	}
}
