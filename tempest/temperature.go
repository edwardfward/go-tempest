package tempest

type TempUnit int

const (
	Celsius TempUnit = iota
	Fahrenheit
	Kelvin
)

// TempReading in unit degrees.
type TempReading struct {
	Reading float64
	Units   TempUnit
}

// NewTempReading creates a new temperature reading.
func NewTempReading(reading float64, units TempUnit) TempReading {
	return TempReading{
		Reading: reading,
		Units:   units,
	}
}

// C returns the temperature reading in degrees Celsius.
func (t TempReading) C() float64 {
	switch t.Units {
	case Celsius:
		return t.Reading
	case Fahrenheit:
		return (t.Reading - 32.0) * 5 / 9
	case Kelvin:
		return t.Reading - 273.15
	}
	return 0
}

// F returns the temperature reading in degrees Fahrenheit.
func (t TempReading) F() float64 {
	switch t.Units {
	case Celsius:
		return (t.Reading * 9 / 5) + 32.0
	case Fahrenheit:
		return t.Reading
	case Kelvin:
		return (t.Reading-273.15)*9/5 + 32.0
	}
	return 0
}

// K returns the temperature reading in Kelvin.
func (t TempReading) K() float64 {
	switch t.Units {
	case Celsius:
		return t.Reading + 273.15
	case Fahrenheit:
		return (t.Reading-32.0)*5/9 + 273.15
	case Kelvin:
		return t.Reading
	}
	return 0
}
