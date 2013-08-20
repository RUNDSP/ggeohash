package ggeohash

import "testing"

//
// Benchmark the tests
//
func Benchmark_Encode(b *testing.B) {
	location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}

	for i := 0; i < b.N; i++ {
		Encode(location, 12)
	}
}

func Benchmark_EncodePrecisions(b *testing.B) {
	location := &PreciseLocation{Latitude: 37.8324, Longitude: 112.5584}

	for i := 0; i < b.N; i++ {
		EncodePrecisions(location, 2, 7)
	}
}

func Benchmark_DecodeBoundBox(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecodeBoundBox("ww8p1r4t8")
	}
}

func Benchmark_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("ww8p1r4t8")
	}
}

func Benchmark_NeighborNorth(b *testing.B) {
	directions := [2]CardialDirections{North, None}

	for i := 0; i < b.N; i++ {
		Neighbor("dqcjq", directions)
	}
}

func Benchmark_NeighborSouthWest(b *testing.B) {
	directions := [2]CardialDirections{South, West}

	for i := 0; i < b.N; i++ {
		Neighbor("dqcjq", directions)
	}
}
