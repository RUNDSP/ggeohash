package ggeohash

import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestPreciseLocationSpecs(t *testing.T) {
	// Setup the suite
	r := gospec.NewRunner()

	// Add new specs here:
	r.AddSpec(PreciseLocationSpecs)

	// Execute the suite
	gospec.MainGoTest(r, t)
}

// Helpers
func PreciseLocationSpecs(c gospec.Context) {

	c.Specify("[PreciseLocation][Encode] Returns GeoHash", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		value := location.Encode(9)
		c.Expect(string(value), gospec.Equals, "ww8p1r4t8")
	})

	c.Specify("[PreciseLocation][Encode] Caches Latitude", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		location.Encode(9)
		c.Expect(location.encodedLatitude, gospec.Equals, location.Latitude)
	})

	c.Specify("[PreciseLocation][Encode] Caches Longitude", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		location.Encode(9)
		c.Expect(location.encodedLongitude, gospec.Equals, location.Longitude)
	})

	c.Specify("[PreciseLocation][Encode] Caches GeoHashes >= 12", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		location.Encode(9)
		c.Expect(len(location.encodedCache), gospec.Equals, 12)

		location.Encode(15)
		c.Expect(len(location.encodedCache), gospec.Equals, 15)

		location.Encode(1)
		c.Expect(len(location.encodedCache), gospec.Equals, 15)
	})

	c.Specify("[PreciseLocation][EncodeRange] Returns array of GeoHashes", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		values := location.EncodeRange(1, 1)
		c.Expect(values, gospec.Satisfies, nil != values)
		c.Expect(len(values), gospec.Equals, 1)
		c.Expect(string(values[0]), gospec.Equals, "w")
	})

	c.Specify("[PreciseLocation][EncodeRange] Returns array of GeoHashes", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		values := location.EncodeRange(9, 9)
		c.Expect(values, gospec.Satisfies, nil != values)
		c.Expect(len(values), gospec.Equals, 1)
		c.Expect(string(values[0]), gospec.Equals, "ww8p1r4t8")
	})

	c.Specify("encodes latitude & longitude as range or precisions", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		values := location.EncodeRange(1, 5)
		c.Expect(values, gospec.Satisfies, nil != values)
		c.Expect(len(values), gospec.Equals, 5)
		c.Expect(string(values[0]), gospec.Equals, "w")
		c.Expect(string(values[1]), gospec.Equals, "ww")
		c.Expect(string(values[2]), gospec.Equals, "ww8")
		c.Expect(string(values[3]), gospec.Equals, "ww8p")
		c.Expect(string(values[4]), gospec.Equals, "ww8p1")
	})
}
