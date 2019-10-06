package main

import (
	"flag"
	"fmt"
	"log"

	client "github.com/michaelmcallister/coord2locality"
)

func main() {
	var lat = flag.Float64("lat", 0.0, "Latitude of the coordinate")
	var lon = flag.Float64("long", 0.0, "Longitude of the coordinate")
	var filename = flag.String("file", "./postcodes.json", "location of the postcode data")
	flag.Parse()

	c := client.New(*filename)
	result, err := c.CoordinatesToLocation(*lat, *lon)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}
