package tempest

type PressureUnit int

const (
	Millibar PressureUnit = iota
	Pascal
	InHg
	hPa
)

// PressureReading represents a pressure reading.
type PressureReading struct {
	pressure float64
	unit     PressureUnit
}

// InHg returns the pressure in inches of mercury.
func (p *PressureReading) InHg() float64 {
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
func (p *PressureReading) Millibar() float64 {
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
func (p *PressureReading) Pascal() float64 {
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
func (p *PressureReading) Hectopascal() float64 {
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
