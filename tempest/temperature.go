package tempest

type TempUnit int

const (
	Celsius TempUnit = iota
	Fahrenheit
	Kelvin
)

// Temp represents a temperature reading.
type Temp struct {
	Reading float64
	Units   TempUnit
}

// NewTemp creates a new temperature reading.
func NewTemp(reading float64, units TempUnit) Temp {
	return Temp{
		Reading: reading,
		Units:   units,
	}
}

// C returns the temperature reading in degrees Celsius.
func (t Temp) C() float64 {
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
func (t Temp) F() float64 {
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
func (t Temp) K() float64 {
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
