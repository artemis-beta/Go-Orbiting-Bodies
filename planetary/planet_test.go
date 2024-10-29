package planetary

import (
	"math"
	"testing"
)

// Test returned cartesian vector between two objects
func TestVectorTo(t *testing.T) {
	p1_vec := CartesianVector{0, 0, 0}
	p2_vec := CartesianVector{1, 2, 3}
	p1 := Planet{"A", p1_vec, 0}
	p2 := Planet{"B", p2_vec, 0}
	separation := p1.VectorTo(&p2)
	expected := CartesianVector{1, 2, 3}
	if math.Abs(separation.X - expected.X) > 1e-8 ||  math.Abs(separation.Y - expected.Y) > 1e-8 || math.Abs(separation.Z - expected.Z) > 1e-8 {
		t.Errorf("Expected vector %v, got %v", expected, separation)
	}
}

// Test conversion from cartesian to spherical polar systems
func TestCartesiantoPolar(t *testing.T) {
	cartesian := CartesianVector{3, 4, 0}
	if cartesian.Radial() != 5 {
		t.Errorf("Expected radial component %v, got %v", 5, cartesian.Radial())
	}
	expected_angle := math.Asin(4.0 / 5.0)
	if math.Abs(cartesian.Phi() - expected_angle) > 1e-5 {
		t.Errorf("Expected phi component %v, got %v", math.Pi / 2.0, cartesian.Phi())
	}
	if cartesian.Theta() != math.Pi / 2.0 {
		t.Errorf("Expected theta component %v, got %v", expected_angle, cartesian.Theta())
	}
}

// Test conversion from spherical polar to cartesian systems
func TestPolarCartesian(t *testing.T) {
	polar := PolarVector{5, math.Pi / 2.0, math.Asin(4.0 / 5.0)}
	if polar.X() != 3 {
		t.Errorf("Expected X component %v, got %v", 3, polar.X())
	}
	if polar.Y() != 4 {
		t.Errorf("Expected Y component %v, got %v", 4, polar.Y())
	}
	if polar.Z() > 1e-5 {
		t.Errorf("Expected Z component %v, got %v", 0, polar.Z())
	}
}

// Test a multistage orbit in pi/4 intervals
func TestRotation(t *testing.T) {
	origin := CartesianVector{0.0, 0.0, 0.0}
	body_pos := CartesianVector{3.0, 0.0, 0.0}
	central_body := Planet{"O", origin, 0}
	other := Planet{"A", body_pos, 0}

	if other.Position.Phi() != 0.0 {
		t.Errorf("Expected start angle component %v, got %v", 0.0, other.Position.Phi())
	}

	for i := 1; i < 9; i++ { 

		other = other.Orbit(&central_body, math.Pi / 4.0)
		t.Logf("%v", other.Position.Phi())

		if math.Abs(other.Position.Phi() - float64(i) * math.Pi / 4.0) > 1e-8 {
			t.Errorf("[%v/8] Expected new angle component %v, got %v", i, float64(i) * math.Pi / 4.0, other.Position.Phi())
		}
	}
}