package ggeohash

import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestNeighborSpecs(t *testing.T) {
	// Setup the suite
	r := gospec.NewRunner()

	// Add new specs here:
	r.AddSpec(NeighborSpecs)

	// Execute the suite
	gospec.MainGoTest(r, t)
}

// Helpers
func NeighborSpecs(c gospec.Context) {
	c.Specify("[Neighbor] North neighbor", func() {
		value := Neighbor("dqcjq", [2]CardialDirections{North, None})
		c.Expect(string(value), gospec.Equals, "dqcjw")
	})

	c.Specify("[Neighbor] South West neighbor", func() {
		value := Neighbor("dqcjq", [2]CardialDirections{South, West})
		c.Expect(string(value), gospec.Equals, "dqcjj")
	})
}
