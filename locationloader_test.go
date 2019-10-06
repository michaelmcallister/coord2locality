package coord2locality // "github.com/michaelmcallister/coord2locality"

import (
	"testing"
)

func TestLoadInvalidLocation(t *testing.T) {
	l := &LocalLoader{Filepath: "invalid"}
	_, err := l.LoadLocations()
	if err == nil {
		t.Fatal("expected err but got nil error")
	}
}

func TestLoadCachedData(t *testing.T) {
	fakeLocation := Location{
		Locality: "fake",
		Lat:      0.0,
		Long:     0.0,
	}
	l := &LocalLoader{
		cache:    []Location{fakeLocation},
		Filepath: "invalid",
	}
	got, err := l.LoadLocations()
	if err != nil {
		t.Fatalf("expected nil error but got: %v", err)
	}

	if got[0].Locality != fakeLocation.Locality {
		t.Fatalf("non-cached result. want: %+v\n got: %+v\n", got[0], fakeLocation)
	}
}

func TestLoadSuccess(t *testing.T) {
	l := &LocalLoader{Filepath: "test/postcodes.json"}
	got, err := l.LoadLocations()
	if err != nil {
		t.Fatalf("expected nil error but got %v", err)
	}
	if got[0].Locality != "PYRMONT" {
		t.Fatalf("unexpected result. got: %+v\n", got[0])
	}
}

func TestLoadFailure(t *testing.T) {
	l := &LocalLoader{Filepath: "test/invalid.txt"}
	_, err := l.LoadLocations()
	if err == nil {
		t.Fatalf("expected non nil error but got nil")
	}
}
