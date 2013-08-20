package ggeohash

import "os"
import "encoding/csv"
import "strconv"
import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestEncodeSpecs(t *testing.T) {
	// Setup the suite
	r := gospec.NewRunner()

	// Add new specs here:
	r.AddSpec(EncodeSpecs)

	// Execute the suite
	gospec.MainGoTest(r, t)
}

// Helpers
func EncodeSpecs(c gospec.Context) {
	c.Specify("[Encode] Encodes latitude & longitude to GeoHash", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		value := Encode(location, 9)
		c.Expect(string(value), gospec.Equals, "ww8p1r4t8")
	})

	c.Specify("encodes latitude & longitude as range or precisions", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		values := EncodePrecisions(location, (1), (1))

		c.Expect(values, gospec.Satisfies, nil != values)
		c.Expect(len(values), gospec.Equals, 1)
		c.Expect(string(values[0]), gospec.Equals, "w")
	})

	c.Specify("encodes latitude & longitude as range or precisions", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		values := EncodePrecisions(location, (9), (9))

		c.Expect(values, gospec.Satisfies, nil != values)
		c.Expect(len(values), gospec.Equals, 1)
		c.Expect(string(values[0]), gospec.Equals, "ww8p1r4t8")
	})

	c.Specify("encodes latitude & longitude as range or precisions", func() {
		location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}
		values := EncodePrecisions(location, 1, 5)

		c.Expect(values, gospec.Satisfies, nil != values)
		c.Expect(len(values), gospec.Equals, 5)
		c.Expect(string(values[0]), gospec.Equals, "w")
		c.Expect(string(values[1]), gospec.Equals, "ww")
		c.Expect(string(values[2]), gospec.Equals, "ww8")
		c.Expect(string(values[3]), gospec.Equals, "ww8p")
		c.Expect(string(values[4]), gospec.Equals, "ww8p1")
	})

	//
	// Batch-test the GeoHash Encode against the Node.js's output.
	//
	// It is critical this is correct and consistent across platforms.
	//
	// There are multiple imlementations of this in several languages:
	// - GO
	// - Node.JS
	// - Ruby
	// - C++
	// - etc
	//
	c.Specify("CSV of encoded latitude, longitude, and precision matches encode", func() {
		file, err := os.Open("./encode_precise_location_test.csv")
		if nil != err {
			panic(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		lines, err := reader.ReadAll()
		if nil != err {
			panic(err)
		}

		for index, line := range lines {
			if index == 0 {
				continue
			}
			i := 0
			latitude, _ := strconv.ParseFloat(line[i], 64)
			i++
			longitude, _ := strconv.ParseFloat(line[i], 64)
			i++
			precision, _ := strconv.Atoi(line[i])
			i++
			expected := line[i]
			i++

			value := Encode(&PreciseLocation{Latitude: latitude, Longitude: longitude}, (precision))
			c.Expect(string(value), gospec.Equals, expected)
		}
	})
}
