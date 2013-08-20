package ggeohash

import "strings"
import "fmt"

//
// Bounded box for parsing latitude and longitude
//
type DecodedBoundBox struct {
	Min, Max *PreciseLocation
}

func MakeDecodedBoundBox() *DecodedBoundBox {
	output := &DecodedBoundBox{}

	// -90, -180
	output.Min = MakePreciseLocationLowerBound()
	// +90, +180
	output.Max = MakePreciseLocationUpperBound()

	return output
}

func DecodeBoundBox(hash_string string) *DecodedBoundBox {
	// Downcase the input string
	hash_string = strings.ToLower(hash_string)

	output := MakeDecodedBoundBox()

	var islon bool = true

	for i, c := range hash_string {
		byte_at := byte(c)
		index_at := convertByteToIndex(byte_at)
		if index_at < 0 {
			msg := fmt.Sprintf("[DecodeBoundBox]  Unknown byte at index=%d, rune='%v' in string='%v'", i, c, hash_string)
			panic(msg)
		}

		index_of := uint8(index_at)
		for bits := 4; bits >= 0; bits-- {
			bit := (index_of >> uint8(bits)) & 1
			if islon {
				mid := (output.Max.Longitude + output.Min.Longitude) / 2
				if bit == 1 {
					output.Min.Longitude = mid
				} else {
					output.Max.Longitude = mid
				}
			} else {
				mid := (output.Max.Latitude + output.Min.Latitude) / 2
				if bit == 1 {
					output.Min.Latitude = mid
				} else {
					output.Max.Latitude = mid
				}
			}
			islon = !islon
		}
	}

	return output
}
