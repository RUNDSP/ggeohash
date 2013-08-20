package ggeohash

const LatitudeLowerBound = -90
const LatitudeUpperBound = 90
const LongitudeLowerBound = -180
const LongitudeUpperBound = 180

type PreciseLocation struct {
	Latitude  float64
	Longitude float64
}

func MakePreciseLocationLowerBound() *PreciseLocation {
	return &PreciseLocation{Latitude: LatitudeLowerBound, Longitude: LongitudeLowerBound}
}

func MakePreciseLocationUpperBound() *PreciseLocation {
	return &PreciseLocation{Latitude: LatitudeUpperBound, Longitude: LongitudeUpperBound}
}
