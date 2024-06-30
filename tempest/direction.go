package tempest

type DirectionUnit int

const (
	Cardinal DirectionUnit = iota
	Degrees
	Radians
	Mils
)

// Direction represents a direction measurement reported
// by a sensor to the hub.
type Direction struct {
	direction interface{}
	unit      DirectionUnit
}

// NewDirection creates a new direction measurement with a
// reading and the unit of measurement.
func NewDirection(reading interface{}, unit DirectionUnit) Direction {
	switch unit {
	case Cardinal:
		switch reading.(type) {
		case string:
			return Direction{
				direction: cardinalDegrees(reading.(string)),
			}
		}
	case Degrees:
		switch reading.(type) {
		case float64:
			return Direction{
				direction: reading.(float64),
				unit:      unit,
			}
		}
	default:
		panic("unhandled default case")
	}

	return Direction{
		direction: reading,
		unit:      unit,
	}
}

// Cardinal returns the direction as a cardinal direction.
func (d Direction) Cardinal() string {
	switch d.unit {
	case Cardinal:
		return cardinalDirection(d.direction)
	case Degrees:
		return cardinalDirection(d.direction)
	case Radians:
		return cardinalDirection(d.direction * 180 / 3.14159)
	case Mils:
		return cardinalDirection(d.direction * 3200 / 6400)
	}
	return ""
}

// Degrees from 0 to 359.9.
func (d Direction) Degrees() float64 {
	switch d.unit {
	case Cardinal:
		return cardinalDegrees(&d.direction)
	case Degrees:
		return d.direction
	case Radians:
		return d.direction * 180 / 3.14159
	case Mils:
		return d.direction * 3200 / 6400
	}
	return 0

}

// cardinalDirection returns the cardinal direction for a
// given degree measurement.
func cardinalDirection(deg float64) string {
	// Normalize the degree to 0-359.9
	deg = deg - 360*float64(int(deg/360))

	switch {
	case deg >= 348.75 || deg < 11.25:
		return "N"
	case deg < 33.75:
		return "NNE"
	case deg < 56.25:
		return "NE"
	case deg < 78.75:
		return "ENE"
	case deg < 101.25:
		return "E"
	case deg < 123.75:
		return "ESE"
	case deg < 146.25:
		return "SE"
	case deg < 168.75:
		return "SSE"
	case deg < 191.25:
		return "S"
	case deg < 213.75:
		return "SSW"
	case deg < 236.25:
		return "SW"
	case deg < 258.75:
		return "WSW"
	case deg < 281.25:
		return "W"
	case deg < 303.75:
		return "WNW"
	case deg < 326.25:
		return "NW"
	case deg < 348.75:
		return "NNW"
	}
	return ""
}

// cardinalDegrees returns the degree measurement for a
// given cardinal direction.
func cardinalDegrees(cardinalDirection string) float64 {
	switch cardinalDirection {
	case "N":
		return 0
	case "NNE":
		return 22.5
	case "NE":
		return 45
	case "ENE":
		return 67.5
	case "E":
		return 90
	case "ESE":
		return 112.5
	case "SE":
		return 135
	case "SSE":
		return 157.5
	case "S":
		return 180.0
	case "SSW":
		return 202.5
	case "SW":
		return 225.0
	case "WSW":
		return 247.5
	case "W":
		return 270.0
	case "WNW":
		return 292.5
	case "NW":
		return 315.0
	case "NNW":
		return 337.5
	default:
		return 0
	}
}
