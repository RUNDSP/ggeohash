package ggeohash

import "math"

const LatitudeLowerBound = -90
const LatitudeUpperBound = 90
const LongitudeLowerBound = -180
const LongitudeUpperBound = 180

type PreciseLocation struct {
	Latitude  float64
	Longitude float64

	// Cached Encoded Geo Hashes
	encodedLatitude, encodedLongitude float64
	encodedCache                      []byte
}

func (p *PreciseLocation) Encode(precision int) []byte {
	// Has the Latitude Changed?
	// Has the Longitude Changed?
	// Is the buffer is too small?
	if len(p.encodedCache) < precision || p.encodedLatitude != p.Latitude || p.encodedLongitude != p.Longitude {
		// Invalidate the cache
		p.encodedCache = nil
	}

	// If we don't have a cached encoding, then create one now
	if nil == p.encodedCache {
		// Save the Latitude and Longitude
		p.encodedLatitude = p.Latitude
		p.encodedLongitude = p.Longitude

		// Encode the geo hash at 12 (or precision if it is larger)
		max := math.Max(float64(12), float64(precision))
		p.encodedCache = Encode(p, int(max))
	}

	// Return the slice with the encoded bytes
	return p.encodedCache[0:precision]
}

func (p *PreciseLocation) EncodeRange(precision_min int, precision_max int) [][]byte {
	// Swap the min/max if the user messed them up
	if precision_max < precision_min {
		precision_min, precision_max = precision_max, precision_min
	}

	// Work backwards from the Max to Min precisions
	length := precision_max - precision_min + 1
	output := make([][]byte, length)

	// Max precision
	output[length-1] = p.Encode(precision_max)

	// Count up to precision_max
	for i := 0; i < length-1; i++ {
		output[i] = output[length-1][0 : precision_min+i]
	}

	return output
}

func MakePreciseLocationLowerBound() *PreciseLocation {
	return &PreciseLocation{Latitude: LatitudeLowerBound, Longitude: LongitudeLowerBound}
}

func MakePreciseLocationUpperBound() *PreciseLocation {
	return &PreciseLocation{Latitude: LatitudeUpperBound, Longitude: LongitudeUpperBound}
}
