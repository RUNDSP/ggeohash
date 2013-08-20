package ggeohash

import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestDecodeBoundBoxSpecs(t *testing.T) {
	// Setup the suite
	r := gospec.NewRunner()

	// Add new specs here:
	r.AddSpec(DecodeBoundBoxSpecs)

	// Execute the suite
	gospec.MainGoTest(r, t)
}

// Helpers
func DecodeBoundBoxSpecs(c gospec.Context) {
	c.Specify("[DecodeBoundBox] Decodes GeoHash to DecodedBoundBox", func() {
		value := DecodeBoundBox("ww8p1r4t8")

		c.Expect(value.Min.Latitude, gospec.Equals, 37.83236503601074)
		c.Expect(value.Min.Longitude, gospec.Equals, 112.55836486816406)
		c.Expect(value.Max.Latitude, gospec.Equals, 37.83240795135498)
		c.Expect(value.Max.Longitude, gospec.Equals, 112.5584077835083)
	})
}
