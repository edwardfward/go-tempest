package tempest

import "time"

// LightningStrikeEvent represents a lightning strike event reported by
// a sensor to the hub.
type LightningStrikeEvent struct {
	Sensor    *WeatherSensor
	Hub       *Hub
	EventTime time.Time
	Distance  float64 // Distance in kilometers from the sensor to the strike.
	Energy    int     // Energy of the strike. There is no unit of measure for this value.
}
