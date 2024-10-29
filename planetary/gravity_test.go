package planetary

import (
	"math"
	"testing"
)

func TestGravitationalForce(t *testing.T) {
	planet_a := Planet{
		"A", CartesianVector{0, 0, 0}, 1e23,
	}
	planet_b := Planet{
		"B", CartesianVector{1.3e7, 0, 0}, 1e12,
	}

	gravity := planet_a.NewtonGravityForce(&planet_b)

	expected := CartesianVector{3.9493E10, 0, 0}

	if math.Abs(gravity.X - expected.X) / expected.X > 1e-5 {
		t.Errorf("Expected force of %vN, got %vN", expected, gravity)
	}
}

func TestKeplerPeriod(t *testing.T) {
	planet_a := Planet{
		"A", CartesianVector{0, 0, 0}, 1e23,
	}
	planet_b := Planet{
		"B", CartesianVector{1.3e7, 0, 0}, 1e12,
	}

	period := planet_a.KeplerPeriod(&planet_b)

	expected := 1.13997e5

	if math.Abs(period - expected) / expected > 1e-5 {
		t.Errorf("Expected period of %vs, got %vs", expected, period)
	}
}