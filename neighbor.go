package ggeohash

type CardialDirections int

const (
	None  CardialDirections = 0
	North                   = 1
	South                   = -1
	East                    = 1
	West                    = -1
)

func Neighbor(hash_string string, direction [2]CardialDirections) []byte {
	// Adjust the DecodedPosition for the direction of the neighbors
	position := Decode(hash_string)
	precision := (len(hash_string))

	location := &PreciseLocation{}
	location.Latitude = position.Location.Latitude + float64(direction[0])*position.Error.Latitude*2
	location.Longitude = position.Location.Longitude + float64(direction[1])*position.Error.Longitude*2

	return Encode(location, precision)
}
