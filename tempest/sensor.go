package tempest

import "time"

// WeatherSensor represents a Tempest weather sensor.
type WeatherSensor struct {
	SensorSerial string     `json:"serial_number"`
	LastMessage  time.Time  `json:"last_message"`
	Coordinate   Coordinate `json:"coordinate"`
}
