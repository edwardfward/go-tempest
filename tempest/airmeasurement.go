package tempest

// AirMeasurement is a sample of air data from a sensor.
type AirMeasurement struct {
	Sensor      *WeatherSensor  // The sensor that sent the data.
	Hub         *Hub            // The hub that received the data.
	EventTime   int64           // Epoch seconds UTC.
	Temperature TempReading     // Air temp reading in degrees Celsius.
	Humidity    float64         // Air temp reading relative humidity percentage (0-100%).
	Pressure    PressureReading // Air pressure reading in millibars (hPa).
}

// C returns the air temperature in degrees Celsius.
func (a *AirMeasurement) C() float64 {
	return a.Temperature.C()
}

// F returns the air temperature in degrees Fahrenheit.
func (a *AirMeasurement) F() float64 {
	return a.Temperature.F()
}

// K returns the air temperature in Kelvin.
func (a *AirMeasurement) K() float64 {
	return a.Temperature.K()
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
	return a.Pressure.Pascal()
}

// InHg returns the pressure in inches of mercury.
func (a *AirMeasurement) InHg() float64 {
	return a.Pressure.InHg()
}
