package tempest

import "time"

type WeatherSensor struct {
	SensorSerialNumber string    `json:"serial_number"`
	LastMessage        time.Time `json:"last_message"`
}
