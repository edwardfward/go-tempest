package tempest

type Type string

// Tempest message types.
const (
	MessageTypeRainStartEvent  = "evt_precip"
	MessageTypeLightningStrike = "evt_strike"
	MessageTypeRapidWind       = "rapid_wind"
	MessageTypeObservation     = "obs_st"
	MessageTypeAirObservation  = "obs_air"
	MessageTypeSkyObservation  = "obs_sky"
	MessageTypeDeviceStatus    = "device_status"
	MessageTypeHubStatus       = "hub_status"
)

// validTypes is a map of valid Tempest message types.
// We use a map here to make it easy and fast to check
// if a message type is valid or not.
var validTypes = map[Type]struct{}{
	MessageTypeRainStartEvent:  {},
	MessageTypeLightningStrike: {},
	MessageTypeRapidWind:       {},
	MessageTypeObservation:     {},
	MessageTypeAirObservation:  {},
	MessageTypeSkyObservation:  {},
	MessageTypeDeviceStatus:    {},
	MessageTypeHubStatus:       {},
}

// Valid returns true if the message type is a valid
// Tempest message type.
func (t Type) Valid() bool {
	_, ok := validTypes[t]
	return ok
}
