package tempest

import (
	"context"
	"time"
)

// HubRepo is an interface for storing and loading Tempest hubs.
type HubRepo interface {
	LoadHubs(ctx context.Context) ([]*Hub, error)
	SaveHub(ctx context.Context, hub *Hub) error
}

// SensorRepo is an interface for storing and loading weather sensors.
type SensorRepo interface {
	LoadSensors() ([]*WeatherSensor, error)
	SaveSensor(*WeatherSensor) error
}

type MessageRepo interface {
	LoadMessages(ctx context.Context, start, end time.Time, msg Type)
	SaveMessage(ctx context.Context, message WeatherMessage)
	DeleteMessages(ctx context.Context, start, end time.Time, msg Type)
}
