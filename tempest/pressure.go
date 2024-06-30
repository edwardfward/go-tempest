package tempest

type PressureUnit int

const (
	Millibar PressureUnit = iota
	Pascal
	InHg
	hPa
)

// Pressure represents a pressure reading.
type Pressure struct {
	pressure float64
	unit     PressureUnit
}

// InHg returns the pressure in inches of mercury.
func (p *Pressure) InHg() float64 {
	switch p.unit {
	case Millibar:
		return p.pressure / 33.8639
	case Pascal:
		return p.pressure / 3386.389
	case hPa:
		return p.pressure / 33.8639
	case InHg:
		return p.pressure
	}

	return 0.0
}

// Millibar returns the pressure in millibars.
func (p *Pressure) Millibar() float64 {
	switch p.unit {
	case Millibar:
		return p.pressure
	case Pascal:
		return p.pressure / 100
	case hPa:
		return p.pressure
	case InHg:
		return p.pressure * 33.8639
	}

	return 0.0
}

// Pascal returns the pressure in pascals.
func (p *Pressure) Pascal() float64 {
	switch p.unit {
	case Millibar:
		return p.pressure * 100
	case Pascal:
		return p.pressure
	case hPa:
		return p.pressure * 100
	case InHg:
		return p.pressure * 3386.389
	}

	return 0.0
}

// Hectopascal returns the pressure in hectopascals.
func (p *Pressure) Hectopascal() float64 {
	switch p.unit {
	case Millibar:
		return p.pressure
	case Pascal:
		return p.pressure / 100
	case hPa:
		return p.pressure
	case InHg:
		return p.pressure * 33.8639
	}

	return 0.0
}

// SeaLevelPressure returns the pressure adjusted for elevation.
func (p *Pressure) SeaLevelPressure(elevation Distance) Pressure {
	// Convert the elevation to meters.
	meters := elevation.Meters()

	// Convert the pressure to pascals.
	pascals := p.Pascal()

	// Calculate the pressure at sea level.
	seaLevel := pascals / (1 - (meters / 44330.0))

	// Return the adjusted pressure.
	return Pressure{
		pressure: seaLevel,
		unit:     Pascal,
	}
}
