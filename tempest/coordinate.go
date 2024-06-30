package tempest

type Latitude float64
type Longitude float64
type Elevation float64

// Coordinate represents a geographic coordinate.
type Coordinate struct {
	Latitude  Latitude
	Longitude Longitude
	Elevation Elevation
}
