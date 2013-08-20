package ggeohash

import "bytes"

// Static array of 0-9, a-z
var base32_codes [32]byte = [32]byte{}

func init() {
	characters := "0123456789bcdefghjkmnpqrstuvwxyz"

	// Map the bytes to bytes & index positions
	for index, rune := range characters {
		byte_at := byte(rune)
		base32_codes[index] = byte_at
	}
}

//
// Convert the given index to it's byte
// Assumes i is within [0, 32)
//
func convertIndexToByte(i int) byte {
	return base32_codes[i]
}

//
// Convert the given byte to an int index
// == -1 --> Failure!
// >= 0  --> Success!
//
func convertByteToIndex(b byte) int {
	// NOTE: This should be thread save
	return bytes.IndexByte(base32_codes[:], b)
}
