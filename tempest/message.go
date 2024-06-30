package tempest

type WeatherMessage interface {
	Type() Type
	Read(rawMessage []byte) error
}
