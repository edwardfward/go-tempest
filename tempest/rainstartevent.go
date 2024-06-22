package tempest

import "time"

// RainStartEvent represents a rain start event.
type RainStartEvent struct {
	SensorSerialNumber string    `json:"serial_number"` // Serial number of the sensor reporting the event to the hub
	Hub                Hub       `json:"hub"`           // The hub reporting the event to the network.
	EventTime          time.Time `json:"evt"`           // Timestamp(Epoch UTC) of the event.
}

// Type returns the message type for the rain start event.
func (r *RainStartEvent) Type() Type {
	return MessageTypeRainStartEvent
}

// Time returns the time of the rain start event.
func (r *RainStartEvent) Time() time.Time {
	return r.EventTime
}

// TimeSince will return the amount of time from the rain start event.
func (r *RainStartEvent) TimeSince() time.Duration {
	return time.Since(r.EventTime)
}
