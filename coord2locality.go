package coord2locality // import "github.com/michaelmcallister/coord2locality"

import (
	"math"
)

// Haversine implementation.
type Haversine struct {
	loader LocationLoader
}

// New returns a new instance of Haversine.
func New(filepath string) *Haversine {
	return &Haversine{
		loader: &LocalLoader{Filepath: filepath},
	}
}

// toRadians takes a float64 in degrees and returns the value in radians.
func toRadians(i float64) float64 {
	return i * math.Pi / 180
}

// haversine is a function that takes two lat/long pairs and returns the
// distance in meters.
// It is adapted from: https://www.movable-type.co.uk/scripts/latlong.html
func (h Haversine) haversine(latP1, lonP1, latP2, lonP2 float64) float64 {
	const earthRadius = 6371
	l1 := toRadians(latP1)
	l2 := toRadians(latP2)

	d1 := toRadians(latP2 - latP1)
	d2 := toRadians(lonP2 - lonP1)

	a := math.Sin(d1/2)*math.Sin(d1/2) + math.Cos(l1)*math.Cos(l2)*math.Sin(d2/2)*math.Sin(d2/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

// CoordinatesToLocation will return the nearest Location given the lat/long.
func (h Haversine) CoordinatesToLocation(lat, lon float64) (Location, error) {
	locations, err := h.loader.LoadLocations()
	if len(locations) == 0 || err != nil {
		return Location{}, err
	}

	// Seed data.
	minDistance := h.haversine(locations[0].Lat, locations[0].Long, lat, lon)
	resultLocation := locations[0]

	for _, location := range locations {
		distance := h.haversine(location.Lat, location.Long, lat, lon)
		if distance < minDistance {
			minDistance = distance
			resultLocation = location
		}
	}
	return resultLocation, nil
}
