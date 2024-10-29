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

	const N_STEPS int = 1e3

	// Define cartesian coordinates for each body
	planet_a_vec := planetary.CartesianVector{X: 0, Y: 0, Z: 0}
	planet_b_vec := planetary.CartesianVector{X: 4e3, Y: 1e3, Z: 0}
	planet_c_vec := planetary.CartesianVector{X: 4.3e3, Y: 1.2e3, Z: 0}
	planet_d_vec := planetary.CartesianVector{X: 4.35e3, Y: 1.27e3, Z: 0}

	// Define the orbiting bodies
	planet_a := planetary.Planet{Name: "A", Position: planet_a_vec, Mass: 1e17}
	planet_b := planetary.Planet{Name: "B", Position: planet_b_vec, Mass: 1e15}
	planet_c := planetary.Planet{Name: "C", Position: planet_c_vec, Mass: 1e13}
	planet_d := planetary.Planet{Name: "D", Position: planet_d_vec, Mass: 1e10}

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
		planet_b_new := planet_b.Orbit(&planet_a)

		// Find the vector translation of planet B
		// and apply it to all orbiting bodies
		b_offset := planet_b.VectorTo(&planet_b_new)
		planet_c.Position = planet_c.Position.Add(&b_offset)
		planet_d.Position = planet_d.Position.Add(&b_offset)
		planet_c_new := planet_c.Orbit(&planet_b_new)

		// Find the vector translation of planet C
		// and apply it to all orbiting bodies
		c_offset := planet_c.VectorTo(&planet_c_new)
		planet_d.Position = planet_d.Position.Add(&c_offset)

		planet_d = planet_d.Orbit(&planet_c_new)
		planet_c = planet_c_new
		planet_b = planet_b_new 

	}
}