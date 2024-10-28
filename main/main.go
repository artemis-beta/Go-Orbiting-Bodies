/*
	Orbiting body demonstration
	----------------------------

	This example demonstrates four bodies orbitting one another.
	Note as a simplest case mass is not yet a factor in the paths of orbits.
*/

package main

import (
	"log"
	"os"
	"planetary"
)

func main() {

	const N_STEPS int = 500

	// Define cartesian coordinates for each body
	planet_a_vec := planetary.CartesianVector{X: 0, Y: 0, Z: 0}
	planet_b_vec := planetary.CartesianVector{X: 4, Y: 1, Z: 0}
	planet_c_vec := planetary.CartesianVector{X: 6, Y: 2, Z: 0}
	planet_d_vec := planetary.CartesianVector{X: 3, Y: 5, Z: 0}

	// Define the orbiting bodies
	planet_a := planetary.Planet{Name: "A", Position: planet_a_vec}
	planet_b := planetary.Planet{Name: "B", Position: planet_b_vec}
	planet_c := planetary.Planet{Name: "C", Position: planet_c_vec}
	planet_d := planetary.Planet{Name: "D", Position: planet_d_vec}

	// Open data file for writing
	f, err := os.Create("data.csv")

	if err != nil {
        log.Fatal(err)
    }

	// Close file on exit
	defer f.Close()

	// For N time steps plot the orbits iterating through each object
	// and writing the cartesian coordinates to a comma separated file
	for i := 0; i < N_STEPS; i++ {
		for _, planet := range []planetary.Planet{planet_a, planet_b, planet_c, planet_d} {
			_, err := f.WriteString(planet.String() + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		planet_d = planet_d.Orbit(&planet_c, 0.16)
		planet_c = planet_c.Orbit(&planet_b, 0.1)
		planet_b = planet_b.Orbit(&planet_a, 0.05)
	}
}