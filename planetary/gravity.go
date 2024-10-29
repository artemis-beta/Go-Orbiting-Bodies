package planetary

import "math"

const G float64 = 6.67430e-11

// Calculate the force due to gravity
func (p *Planet) NewtonGravityForce(o *Planet) CartesianVector {
	separation := p.VectorTo(o)
	magnitude := G * p.Mass * o.Mass * math.Pow(separation.Radial(), -2.0)
	unit := separation.Unit()
	return unit.Scale(magnitude)
}

// Calculate the orbital time period
func (p *Planet) KeplerPeriod(o *Planet) float64 {
	separation := p.VectorTo(o)

	// The mass used for time period is the larger of the two
	// this function does not care about ordering of the
	// two planet instances
	mass := p.Mass

	if mass < o.Mass {
		mass = o.Mass
	}

	inner_ratio := math.Pow(separation.Radial(), 3.0) / (G * mass)
	return 2 * math.Pi * math.Sqrt(inner_ratio)
}

// Calculate the orbital velocity
func (p *Planet) KeplerAngularVelocity(o *Planet) float64 {
	return 2 * math.Pi / p.KeplerPeriod(o)
}