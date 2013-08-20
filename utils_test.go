package ggeohash

import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestUtilsSpecs(t *testing.T) {
	// Setup the suite
	r := gospec.NewRunner()

	// Add new specs here:
	r.AddSpec(UtilsSpecs)

	// Execute the suite
	gospec.MainGoTest(r, t)
}

// Helpers
func UtilsSpecs(c gospec.Context) {
	//
	// Verify the conversion methods return the correct input/output
	//
	c.Specify("Converts ints to bytes and back again", func() {
		// This is important to get right ...
		// The GeoHash algorithm depends on this being fast and consistent
		src := "0123456789bcdefghjkmnpqrstuvwxyz"
		for i := 0; i < 32; i++ {
			// What is the character at "i"
			src_at := byte(src[i])

			// Map the position to the byte
			byte_at := convertIndexToByte(i)
			c.Expect(byte_at, gospec.Equals, src_at)

			// Map the byte back to the position
			index_at := convertByteToIndex(byte_at)
			c.Expect(index_at, gospec.Equals, i)
		}
	})
}
