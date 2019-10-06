package coord2locality // "github.com/michaelmcallister/coord2locality"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// LocationLoader defines interface necessary to load location data.
type LocationLoader interface {
	LoadLocations() ([]Location, error)
}

// Location contains lat/long co-ords for the respective locality.
type Location struct {
	Postcode string
	Locality string
	State    string
	Lat      float64
	Long     float64
}

// LocalLoader represents local.
type LocalLoader struct {
	cache    []Location
	Filepath string
}

// LoadLocations will read the JSON file from disk and into memory.
func (l LocalLoader) LoadLocations() ([]Location, error) {
	if len(l.cache) > 0 {
		return l.cache, nil
	}
	var locations []Location

	f, err := os.Open(l.Filepath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("Error reading bytes: %v", err)
	}

	err = json.Unmarshal(b, &locations)
	l.cache = locations
	return locations, err
}
