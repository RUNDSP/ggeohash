package ggeohash

import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestDecodePositionSpecs(t *testing.T) {
	// Setup the suite
	r := gospec.NewRunner()

	// Add new specs here:
	r.AddSpec(DecodePositionSpecs)

	// Execute the suite
	gospec.MainGoTest(r, t)
}

// Helpers
func DecodePositionSpecs(c gospec.Context) {
	c.Specify("[Decode] Decodes GeoHash to PreciseLocation + Error", func() {
		value := Decode("ww8p1r4t8")

		c.Expect(value.Location.Latitude, gospec.Satisfies, value.Location.Latitude >= 37.83 && value.Location.Latitude <= (37.84))
		c.Expect(value.Location.Longitude, gospec.Satisfies, value.Location.Longitude >= 112.55 && value.Location.Longitude <= (112.56))
	})
}
