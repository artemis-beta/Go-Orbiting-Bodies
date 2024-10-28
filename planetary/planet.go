// Basic example package for trialling the Go language, currently this behaves rather like
// a spirograph showing the orbits of objects with respect to each other. Currently the
// masses of the objects are not taken into account for simplicity.
//
// Uses conversion between cartesian and spherical polar coordinate systems to write
// the paths of these orbits to stdout

package planetary

import (
	"fmt"
	"math"
)

// Cartesian coordinate object
type CartesianVector struct {
	X float64
	Y float64
	Z float64
}

// Spherical Polar coordinate object
type PolarVector struct {
	R float64
	Theta float64
	Phi float64
}

// Return a printable form of the cartesian vector
func (v* CartesianVector) String() string {
	return fmt.Sprintf("%v, %v, %v", v.X, v.Y, v.Z)
}

// Return a printable form of the polar vector
func (v* PolarVector) String() string {
	return fmt.Sprintf("%v, %v, %v", v.R, v.Theta, v.Phi)
}

// Radial distance calculated as the magnitude of the cartesian vector
func (v *CartesianVector) Radial() float64 {
	return math.Sqrt(math.Pow(v.X, 2.0) + math.Pow(v.Y, 2.0) + math.Pow(v.Z, 2.0))
}

// Polar angle from a cartesian coordinate
func (v *CartesianVector) Theta() float64 {
	
	// Return an angle of zero if the radial
	// component is 0 as opposed to undefined
	if v.Radial() == 0.0 {
		return 0.0
	}

	return math.Acos(v.Z / v.Radial())
}

// Azimuthal angle from a cartesian coordinate
func (v *CartesianVector) Phi() float64 {
	// Azimuthal Angle
	divisor := math.Sqrt(math.Pow(v.X, 2.0) + math.Pow(v.Y, 2.0))
	if divisor == 0.0 {
		return 0.0
	}
	angle := math.Atan2(v.Y, v.X)

	// Ensure angles range from [0, 2pi]
	// as opposed to [-pi, pi]
	if angle < 0 { angle += 2 * math.Pi }

	return angle
}

// The sum of two cartesian vectors
func (v CartesianVector) Add(o *CartesianVector) CartesianVector {
	return CartesianVector{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// The sum of two spherical polar vectors
func (v PolarVector) Add(o *PolarVector) PolarVector {
	return PolarVector{v.R + o.R, v.Theta + o.Theta, v.Phi + o.Phi}
}

// Cartesian X from a polar vector
func (v *PolarVector) X() float64 {
	return v.R * math.Sin(v.Theta) * math.Cos(v.Phi)
}

// Cartesian Y from a polar vector
func (v *PolarVector) Y() float64 {
	return v.R * math.Sin(v.Theta) * math.Sin(v.Phi)
}

// Cartesian Z from a polar vector
func (v *PolarVector) Z() float64 {
	return v.R * math.Cos(v.Theta)
}

// Spherical polar form of a cartesian vector
func (v *CartesianVector) Polar() PolarVector {
	return PolarVector{v.Radial(), v.Theta(), v.Phi()}
}

// Cartesian form of a spherical polar vector
func (v *PolarVector) Cartesian() CartesianVector {
	return CartesianVector{v.X(), v.Y(), v.Z()}
}

// Simple orbiting body object
type Planet struct {
	Name string
	Position CartesianVector
}

func (p *Planet) VectorTo(o *Planet) CartesianVector {
	return CartesianVector{
		o.Position.X - p.Position.X,
		o.Position.Y - p.Position.Y,
		o.Position.Z - p.Position.Z,
	}
}

// Given a body 'p' return a new position after 1 second
// due to rotation in the X-Y plane with respect to a
// second body 'o'
func (p *Planet) Orbit(o *Planet, angular_velocity float64) Planet {
	separation := o.VectorTo(p)
	new_separation_polar := separation.Polar()
	new_separation_polar.Phi += angular_velocity
	new_separation := new_separation_polar.Cartesian()

	position := o.Position.Add(&new_separation)
	
	return Planet{
		p.Name,
		position,
	}
}

// Printable form of the orbiting body
func (p* Planet) String() string {
	return fmt.Sprintf("%v, %v", p.Name, p.Position.String())
}