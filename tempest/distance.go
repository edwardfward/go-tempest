package tempest

type DistanceUnit int

const (
	Meters DistanceUnit = iota
	Kilometers
	Miles
	NauticalMiles
	Feet
	Yards
)

// Distance represents a distance measurement.
type Distance struct {
	Distance float64
	Unit     DistanceUnit
}

// NewDistance creates a new distance measurement.
func NewDistance(distance float64, unit DistanceUnit) *Distance {
	return &Distance{Distance: distance, Unit: unit}
}

// Meters the distance in meters.
func (d *Distance) Meters() float64 {
	switch d.Unit {
	case Meters:
		return d.Distance
	case Kilometers:
		return d.Distance * 1000
	case Miles:
		return d.Distance * 1609.34
	case NauticalMiles:
		return d.Distance * 1852
	case Feet:
		return d.Distance * 0.3048
	case Yards:
		return d.Distance * 0.9144
	}
	return 0
}

// Kilometers the distance in kilometers.
func (d *Distance) Kilometers() float64 {
	switch d.Unit {
	case Meters:
		return d.Distance / 1000
	case Kilometers:
		return d.Distance
	case Miles:
		return d.Distance * 1.60934
	case NauticalMiles:
		return d.Distance * 1.852
	case Feet:
		return d.Distance * 0.0003048
	case Yards:
		return d.Distance * 0.0009144
	}
	return 0
}

// Miles the distance in miles.
func (d *Distance) Miles() float64 {
	switch d.Unit {
	case Meters:
		return d.Distance / 1609.34
	case Kilometers:
		return d.Distance / 1.60934
	case Miles:
		return d.Distance
	case NauticalMiles:
		return d.Distance / 1.15078
	case Feet:
		return d.Distance / 5280
	case Yards:
		return d.Distance / 1760
	}
	return 0
}

// NauticalMiles the distance in nautical miles.
func (d *Distance) NauticalMiles() float64 {
	switch d.Unit {
	case Meters:
		return d.Distance / 1852
	case Kilometers:
		return d.Distance / 1.852
	case Miles:
		return d.Distance * 1.15078
	case NauticalMiles:
		return d.Distance
	case Feet:
		return d.Distance / 6076.12
	case Yards:
		return d.Distance / 2025.37
	}
	return 0
}

// Feet the distance in feet.
func (d *Distance) Feet() float64 {
	switch d.Unit {
	case Meters:
		return d.Distance / 0.3048
	case Kilometers:
		return d.Distance / 0.0003048
	case Miles:
		return d.Distance * 5280
	case NauticalMiles:
		return d.Distance * 6076.12
	case Feet:
		return d.Distance
	case Yards:
		return d.Distance * 3
	}
	return 0
}

// Yards the distance in yards.
func (d *Distance) Yards() float64 {
	switch d.Unit {
	case Meters:
		return d.Distance / 0.9144
	case Kilometers:
		return d.Distance / 0.0009144
	case Miles:
		return d.Distance * 1760
	case NauticalMiles:
		return d.Distance * 2025.37
	case Feet:
		return d.Distance / 3
	case Yards:
		return d.Distance
	}
	return 0
}
