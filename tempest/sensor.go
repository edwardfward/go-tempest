package tempest

import "time"

// WeatherSensor represents a Tempest weather sensor.
type WeatherSensor struct {
	SensorSerial string    `json:"serial_number"`
	LastMessage  time.Time `json:"last_message"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	E
}
