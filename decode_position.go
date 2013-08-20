package ggeohash

//
// Decoded Latitude,Longitude + error position
//
type DecodedPosition struct {
	Location PreciseLocation

	Error PreciseLocation
}

func Decode(hash_string string) *DecodedPosition {
	bbox := DecodeBoundBox(hash_string)
	output := &DecodedPosition{}
	// Mid point of box
	output.Location.Latitude = (bbox.Min.Latitude + bbox.Max.Latitude) / 2
	output.Location.Longitude = (bbox.Min.Longitude + bbox.Max.Longitude) / 2

	// Mid Point -  Min/Max ==> Error
	output.Error.Latitude = bbox.Max.Latitude - output.Location.Latitude
	output.Error.Longitude = bbox.Max.Longitude - output.Location.Longitude

	return output
}
