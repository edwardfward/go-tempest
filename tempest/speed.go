package tempest

// SpeedUnit represents a unit of speed.
type SpeedUnit int

const (
	MetersPerSecond SpeedUnit = iota
	KilometersPerHour
	MilesPerHour
	Knots
	FeetPerSecond
)

// Speed represents a speed reading from the Tempest device.
type Speed struct {
	speed float64
	unit  SpeedUnit
}

// NewSpeed creates a new speed reading.
func NewSpeed(speed float64, unit SpeedUnit) Speed {
	return Speed{
		speed: speed,
		unit:  unit,
	}
}

// MetersPerSecond converts the speed to meters per second.
func (s *Speed) MetersPerSecond() float64 {
	switch s.unit {
	case MetersPerSecond:
		return s.speed
	case KilometersPerHour:
		return s.speed / 3.6
	case MilesPerHour:
		return s.speed / 2.23694
	case Knots:
		return s.speed / 1.94384
	case FeetPerSecond:
		return s.speed / 3.28084
	}

	return 0
}

// KPH converts the speed to kilometers per hour.
func (s *Speed) KPH() float64 {
	switch s.unit {
	case MetersPerSecond:
		return s.speed * 3.6
	case KilometersPerHour:
		return s.speed
	case MilesPerHour:
		return s.speed * 1.60934
	case Knots:
		return s.speed * 1.852
	case FeetPerSecond:
		return s.speed * 1.09728
	}
	return 0
}

// MPH converts the speed to miles per hour.
func (s *Speed) MPH() float64 {
	switch s.unit {
	case MetersPerSecond:
		return s.speed * 2.23694
	case KilometersPerHour:
		return s.speed * 0.621371
	case MilesPerHour:
		return s.speed
	case Knots:
		return s.speed * 1.15078
	case FeetPerSecond:
		return s.speed * 1.46667
	}
	return 0
}

// KTS converts the speed to knots.
func (s *Speed) KTS() float64 {
	switch s.unit {
	case MetersPerSecond:
		return s.speed * 1.94384
	case KilometersPerHour:
		return s.speed * 0.539957
	case MilesPerHour:
		return s.speed * 0.868976
	case Knots:
		return s.speed
	case FeetPerSecond:
		return s.speed * 0.592484
	}
	return 0
}

// FPS converts the speed to feet per second.
func (s *Speed) FPS() float64 {
	switch s.unit {
	case MetersPerSecond:
		return s.speed * 3.28084
	case KilometersPerHour:
		return s.speed * 0.911344
	case MilesPerHour:
		return s.speed * 1.46667
	case Knots:
		return s.speed * 1.68781
	case FeetPerSecond:
		return s.speed
	}
	return 0
}
