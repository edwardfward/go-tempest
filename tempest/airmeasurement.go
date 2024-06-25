package tempest

// AirMeasurement is a sample of air data from a sensor.
type AirMeasurement struct {
	Sensor      *WeatherSensor
	Hub         *Hub
	EventTime   int64
	Temperature TempReading // degrees Celsius.
	Humidity    float64     // relative humidity percentage.
	Pressure    float64     // millibars.
}

// C returns the air temperature in degrees Celsius.
func (a *AirMeasurement) C() float64 {
	return a.Temperature.C()
}

// F returns the air temperature in degrees Fahrenheit.
func (a *AirMeasurement) F() float64 {
	return (a.Temperature.F() * 9 / 5) + 32.0
}

// K returns the air temperature in Kelvin.
func (a *AirMeasurement) K() float64 {
	return a.Temperature.K() + 273.15
}

// DewPointC returns the dew point in degrees Celsius.
func (a *AirMeasurement) DewPointC() float64 {
	return a.Temperature.C() - ((100.0 - a.Humidity) / 5.0)
}

// DewPointF returns the dew point in degrees Fahrenheit.
func (a *AirMeasurement) DewPointF() float64 {
	return (a.DewPointC() * 9 / 5) + 32.0
}

// Pascal returns the millibar pressure in pascals.
func (a *AirMeasurement) Pascal() float64 {
	return a.Pressure * 100
}

// InHg returns the pressure in inches of mercury.
func (a *AirMeasurement) InHg() float64 {
	return a.Pascal() / 3386.389
}
