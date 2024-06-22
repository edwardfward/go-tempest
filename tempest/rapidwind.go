package tempest

import (
	"fmt"
	"time"
)

const (
	rapidWindEventObservation = "ob"
)

// RapidWindEvent represents a rapid wind event reported by a sensor to the hub.
type RapidWindEvent struct {
	Sensor        *WeatherSensor // Sensor reporting the event to the hub.
	Hub           *Hub           // The hub reporting the event to the network.
	EventTime     time.Time      // Timestamp(Epoch UTC) of the event.
	WindSpeed     float64        // Wind speed in meters per second.
	WindDirection int16          // Wind direction in degrees or where the wind is coming from.
}

// NewRapidWindEvent creates a new rapid wind event with a
// rapid wind raw message from the hub.
func NewRapidWindEvent(raw RawMessage, sensor *WeatherSensor, hub *Hub) (*RapidWindEvent, error) {
	event := &RapidWindEvent{
		Sensor: sensor,
		Hub:    hub,
	}

	rawOb, found := raw[rapidWindEventObservation]
	if !found {
		return nil, fmt.Errorf("rapid wind event observation not found")
	}

	ob, ok := rawOb.([]interface{})
	if !ok {
		return nil, fmt.Errorf("rapid wind event observation is not a float64 slice")
	}

	if len(ob) != 3 {
		return nil, fmt.Errorf("rapid wind event observation does not have 3 elements")
	}

	// The first element is the timestamp of the event.
	if ts, ok := ob[0].(float64); ok {
		event.EventTime = time.Unix(int64(ts), 0)
	} else {
		return nil, fmt.Errorf("rapid wind event observation timestamp cannot be converted to a float64")
	}

	// The second element is the wind speed.
	if speed, ok := ob[1].(float64); ok {
		event.WindSpeed = speed
	} else {
		return nil, fmt.Errorf("rapid wind event observation wind speed is not a float64")
	}

	// The third element is the wind direction.
	if direction, ok := ob[2].(float64); ok {
		event.WindDirection = int16(direction)
	} else {
		return nil, fmt.Errorf("rapid wind event observation wind direction is not a float64")
	}

	return event, nil
}
