package tempest

import "time"

// SensorRepo is an interface for storing and loading weather sensors.
type SensorRepo interface {
	LoadSensors() ([]*WeatherSensor, error)
	SaveSensor(*WeatherSensor) error
}

type WeatherSensor struct {
	SensorSerialNumber string    `json:"serial_number"`
	LastMessage        time.Time `json:"last_message"`
}
