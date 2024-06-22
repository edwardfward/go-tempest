package tempest

// AirMeasurement is a sample of air data from a sensor.
type AirMeasurement struct {
	Sensor      *WeatherSensor
	Hub         *Hub
	EventTime   int64
	Temperature float64 // degrees Celsius.
	Humidity    float64 // relative humidity percentage.
	Pressure    float64 // millibars.
}

// F returns the temperature in degrees Fahrenheit.
func (a *AirMeasurement) F() float64 {
	return (a.Temperature * 9 / 5) + 32.0
}

// K returns the temperature in Kelvin.
func (a *AirMeasurement) K() float64 {
	return a.Temperature + 273.15
}

// DewPointC returns the dew point in degrees Celsius.
func (a *AirMeasurement) DewPointC() float64 {
	return a.Temperature - ((100.0 - a.Humidity) / 5.0)
}

// DewPointF returns the dew point in degrees Fahrenheit.
func (a *AirMeasurement) DewPointF() float64 {
	return (a.DewPointC() * 9 / 5) + 32.0
}

// InHg returns the pressure in inches of mercury.
func (a *AirMeasurement) InHg() float64 {
	return a.Pressure * 0.02953
}