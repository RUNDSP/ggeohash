package main

import "fmt"
import "os"
import "strings"
import ggeohash "../"

import flags "github.com/jessevdk/go-flags"

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	//
	// Operation To Perform
	//
	Encode bool `long:"encode" description:"Encode a geohash"`

	Decode bool `long:"decode" description:"Decode a geohash"`

	Neighbor bool `long:"neighbor" description:"Find the neighbor of a geohash in a given direction: North, South, East, West, North+East, South+West, etc"`

	//
	// Input Parameters for above
	//

	Latitude float64 `long:"latitude" description:"Latitude [-90.0 ... +90.0]"`

	Longitude float64 `long:"longitude" description:"Longitude [-180.0 ... +180.0]"`

	Precision int `long:"precision" description:"Precision [1 ... 12]"`

	GeoHash string `long:"geohash" description:"GeoHash string"`

	North bool `long:"north" description:"Neighbor to the North"`
	South bool `long:"south" description:"Neighbor to the South"`
	East  bool `long:"east" description:"Neighbor to the East"`
	West  bool `long:"west" description:"Neighbor to the West"`
}

func init() {
}

func main() {
	// Parse flags from `args'. Note that here we use flags.ParseArgs for
	// the sake of making a working example. Normally, you would simply use
	// flags.Parse(&opts) which uses os.Args
	args, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(0)
		return
	}

	if opts.Verbose {
		fmt.Printf("Verbosity: %v\n", opts.Verbose)
		fmt.Printf("Encode: %v\n", opts.Encode)
		fmt.Printf("Decode: %v\n", opts.Decode)
		fmt.Printf("Neighbor: %v\n", opts.Neighbor)
		fmt.Printf("Latitude: %v\n", opts.Latitude)
		fmt.Printf("Longitude: %v\n", opts.Longitude)
		fmt.Printf("Precision: %v\n", opts.Precision)
		fmt.Printf("GeoHash: '%v'\n", opts.GeoHash)
		fmt.Printf("Remaining args: [%s]\n", strings.Join(args, " "))
	}

	// Shorthand to make the code below readable
	location := ggeohash.MakePreciseLocationLowerBound()
	location.Latitude = opts.Latitude
	location.Longitude = opts.Longitude

	precision := (opts.Precision)
	geo := opts.GeoHash

	if opts.Encode {
		output := ggeohash.Encode(location, precision)
		fmt.Printf("latitude = %2.10f, longitude = %3.10f, precision = %d, geohash = %s\n", location.Latitude, location.Longitude, precision, output)
	}

	if opts.Decode {
		output := ggeohash.Decode(geo)
		fmt.Printf("geohash = %s, latitude = %2.10f, longitude = %3.10f, latitude.err = %2.10f, longitude.err = %3.10f\n", geo, output.Location.Latitude, output.Location.Longitude, output.Error.Latitude, output.Error.Longitude)
	}

	if opts.Neighbor {
		directions := [2]ggeohash.CardialDirections{ggeohash.None, ggeohash.None}
		if opts.North {
			directions[0] = ggeohash.North
		}
		if opts.South {
			directions[0] = ggeohash.South
		}
		if opts.East {
			directions[1] = ggeohash.East
		}
		if opts.West {
			directions[1] = ggeohash.West
		}

		output := ggeohash.Neighbor(geo, directions)
		fmt.Printf("geohash = %s, directions[0] = %d, directions[1] = %d, neighbor = %s\n", geo, directions[0], directions[1], string(output))
	}
}
