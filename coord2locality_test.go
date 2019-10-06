package coord2locality // "github.com/michaelmcallister/coord2locality"

import (
	"errors"
	"reflect"
	"testing"
)

type fakeLoader struct {
	out []Location
	err error
}

func (f fakeLoader) LoadLocations() ([]Location, error) {
	return f.out, f.err
}

func TestCoordinatesToLocation(t *testing.T) {
	type test struct {
		lat       float64
		lon       float64
		locations []Location
		want      Location
		err       error
	}

	tests := []test{
		{
			lat:       0,
			lon:       0,
			locations: nil,
			want:      Location{},
			err:       errors.New("filesystem error"),
		},
		{
			lat:       0,
			lon:       0,
			locations: nil,
			want:      Location{},
		},
		{
			lat:       0,
			lon:       0,
			want:      Location{Locality: "fake", Lat: 0, Long: 0},
			locations: []Location{Location{Locality: "fake", Lat: 0, Long: 0}},
		},
		{
			lat:  -33.8681091,
			lon:  151.1931297,
			want: Location{Postcode: "2009", State: "NSW", Locality: "PYRMONT", Lat: -33.871222, Long: 151.193055},
			locations: []Location{
				Location{
					Locality: "CANBERRA",
					Postcode: "2600",
					Long:     0,
					Lat:      0,
				},
				Location{
					Locality: "PYRMONT",
					State:    "NSW",
					Postcode: "2009",
					Long:     151.193055,
					Lat:      -33.871222,
				}, Location{
					Locality: "GREENFIELDS",
					State:    "WA",
					Postcode: "6210",
					Long:     115.728286,
					Lat:      -32.557981,
				}},
		},
	}

	for i, tc := range tests {
		fl := fakeLoader{out: tc.locations, err: tc.err}
		haversine := &Haversine{loader: fl}
		got, err := haversine.CoordinatesToLocation(tc.lat, tc.lon)

		if err == nil && tc.err != nil {
			t.Errorf("Unexpected error for #%d. Want nil error but got %v", i, err)
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Unexpected result for #%d. got=%+v\n, want=%+v\n", i, got, tc.want)
		}
	}
}
