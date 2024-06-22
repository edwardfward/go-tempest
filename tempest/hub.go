package tempest

import (
	"net"
	"time"
)

// Hub represents a Tempest hub connected to the network reporting data
// from zero or more sensors connected to the hub.
type Hub struct {
	// Serial number of the hub.
	HubSerialNumber string `json:"hub_sn"`

	// IP address of the hub on the network.
	IPAddress net.IP `json:"ip_address"`

	// Map of sensor serial numbers associated with the hub.
	WeatherSensors map[string]*WeatherSensor `json:"sensors"`

	// Firmware version of the hub.
	FirmwareVersion string `json:"firmware_revision"`

	// Last time the hub was seen on the network.
	LastReported time.Time `json:"report_time"`
}