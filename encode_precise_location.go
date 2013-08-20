package ggeohash

//
// Encode a Latitude and Longitude as a string
//
// Arguments:
//  latitude  float64
//  longitude float64
//  precision int  (ie how long should the hash string be?)
//

func Encode(location *PreciseLocation, precision int) []byte {
	output := make([]byte, precision)
	i := 0

	// DecodedBoundBox for the lat/lon + errors
	bbox := MakeDecodedBoundBox()

	var min_max_avg float64 = 0
	var islon bool = true
	var num_bits uint = 0
	var hash_index int = 0

	for i < precision {
		if islon {
			min_max_avg = (bbox.Max.Longitude + bbox.Min.Longitude) / 2
			if location.Longitude > min_max_avg {
				hash_index = (hash_index << 1) + 1
				bbox.Min.Longitude = min_max_avg
			} else {
				hash_index = (hash_index << 1) + 0
				bbox.Max.Longitude = min_max_avg
			}
		} else {
			min_max_avg = (bbox.Max.Latitude + bbox.Min.Latitude) / 2
			if location.Latitude > min_max_avg {
				hash_index = (hash_index << 1) + 1
				bbox.Min.Latitude = min_max_avg
			} else {
				hash_index = (hash_index << 1) + 0
				bbox.Max.Latitude = min_max_avg
			}
		}
		islon = !islon

		num_bits++
		if 5 == num_bits {
			// Append the byte to the output
			output[i] = (convertIndexToByte(hash_index))
			i++

			// Reset the state counters
			num_bits = 0
			hash_index = 0
		}
	}

	return output
}

//
// Encode a range of GeoHashes. Min and Max are INCLUSIVE!
//
func EncodePrecisions(location *PreciseLocation, precision_min int, precision_max int) [][]byte {
	// Swap the min/max if the user messed them up
	if precision_max < precision_min {
		tmp := precision_min
		precision_min = precision_max
		precision_max = tmp
	}

	// Max precision
	encoded := []byte(Encode(location, precision_max))

	// Work backwards from the Max to Min precisions
	length := precision_max - precision_min + 1
	output := make([][]byte, length)

	// Extract a slice of each precision
	// This doesn't copy the inderlying slice, just aliases it

	// NOTE: Array index and Slice size are counting in opposite directions:
	for i := length - 1; i >= 0; i-- {
		distance_from_end := length - 1 - i

		output[i] = encoded[0 : len(encoded)-distance_from_end]
	}

	return output
}
