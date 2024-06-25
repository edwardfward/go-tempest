package tempest

import "time"

// LightningStrikeEvent represents a lightning strike event reported by
// a sensor to the hub.
type LightningStrikeEvent struct {
	Sensor    *WeatherSensor      // The sensor that detected the strike.
	Hub       *Hub                // The hub that received the data.
	EventTime time.Time           // Epoch seconds UTC.
	Distance  DistanceMeasurement // Distance in kilometers from the sensor to the strike.
	Energy    int                 // Energy of the strike. There is no unit of measure for this value.
}
