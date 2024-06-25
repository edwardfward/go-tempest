package tempest

import "fmt"

const (
	hubSerial    = "hub_sn"
	messageType  = "type"
	sensorSerial = "serial_number"
)

// RawMessage is used to represent a raw message from
// the Tempest hub. It is a map of string keys to
// interface{} values which will be converted to
// the appropriate message type.
type RawMessage map[string]interface{}

// HubSerial returns the raw message's hub serial number.
func (m RawMessage) HubSerial() (string, error) {
	if hub, found := m[hubSerial]; found {
		if serial, ok := hub.(string); ok {
			return serial, nil
		}

		return "", fmt.Errorf("hub serial is not a string")
	}

	// A hub status message will have the hub serial number under the key
	// "serial_number".
	if msgType, err := m.Type(); err == nil && msgType == MessageTypeHubStatus {
		if hub, found := m[sensorSerial]; found {
			if serial, ok := hub.(string); ok {
				return serial, nil
			}

			return "", fmt.Errorf("hub serial is not a string")
		}
	}

	return "", fmt.Errorf("hub serial not found")
}

// SensorSerial returns the raw message's weather sensor serial number.
func (m RawMessage) SensorSerial() (string, error) {
	// A hub status message will not have a sensor serial number.
	if msgType, err := m.Type(); err == nil && msgType == MessageTypeHubStatus {
		return "", fmt.Errorf("weather sensor serial not available in hub status message")
	} else if err != nil {
		return "", fmt.Errorf("error getting message type: %v", err)
	}

	// Valid message type that is not a hub status message.
	if sensor, found := m[sensorSerial]; found {
		if serial, ok := sensor.(string); ok {
			return serial, nil
		}

		return "", fmt.Errorf("sensor serial is not a string")
	}

	return "", fmt.Errorf("sensor serial not found")
}

// Type returns the raw message's message type.
// An error is returned if the message type is not
// present or a valid Tempest message type.
func (m RawMessage) Type() (Type, error) {
	// Check if the message type is present.
	if rawType, ok := m[messageType]; ok {
		if t, ok := rawType.(string); ok {
			msgType := Type(t)

			if msgType.Valid() {
				// Valid message type.
				return msgType, nil
			}

			// Not a valid message type.
			return "", fmt.Errorf("message type invalid: %s", msgType)
		}

		// Cannot convert the message type.
		return "", fmt.Errorf("cannot conver message type")
	}

	// Message type not found.
	return "", fmt.Errorf("message type not found")
}
