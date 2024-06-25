package tempest

// WindMeasurement represents a wind measurement reported by a sensor to the hub.
type WindMeasurement struct {
	Sensor    *WeatherSensor
	Hub       *Hub
	Direction float64      // degrees 0 - 359.9
	Speed     SpeedReading // meters per second
	Time      int64        // Epoch UTC
	Lull      float64      // meters per second
	Gust      float64      // meters per second
}

// KPH wind speed in kilometers per hour.
func (w *WindMeasurement) KPH() float64 {
	return w.Speed.KPH()
}

// KTS wind speed in knots.
func (w *WindMeasurement) KTS() float64 {
	return w.Speed.KTS()
}

// MPH wind speed in miles per hour.
func (w *WindMeasurement) MPH() float64 {
	return w.Speed.MPH()
}

// FPS wind speed in feet per second.
func (w *WindMeasurement) FPS() float64 {
	return w.Speed.FPS()
}
