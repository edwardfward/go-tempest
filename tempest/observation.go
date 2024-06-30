package tempest

const (
	obsIndexTimestampEpochUTC int = iota
	obsIndexWindLull
	obsIndexWindAverage
	obsIndexWindGust
	obsIndexWindDirection
	obsIndexWindSampleInterval
	obsIndexStationPressure
	obsIndexAirTemperature
	obsIndexRelativeHumidity
	obsIndexIlluminance
	obsIndexUV
	obsIndexSolarRadiation // W/m2
	obsIndexRainAccumulationPastMinute
	obsIndexPrecipitationType
	obsIndexLightningStrikeAverageDistance
	obsIndexLightingStrikeCount
	obsIndexBatteryVolts
	obsIndexReportingInterval
)

const (
	EventTypeObservation = "obs_st"
)

// Observation represents a weather observation from a sensor.
type Observation struct {
	SensorSerial string
	HubSerial    string
	Coordinate   Coordinate
	Observations []WeatherObservation
}

// WeatherObservation message type.
type WeatherObservation struct {
	EpochSecondsUTC    int64     // Epoch seconds UTC
	WindLull           Speed     // meters per second
	WindAverage        Speed     // meters per second
	WindGust           Speed     // meters per second
	WindDirection      Direction // degrees
	WindSampleInterval int       // seconds
	StationPressure    Pressure  // millibars
	AirTemperature     Temp      // degrees Celsius
	RelativeHumidity   float64   // percentage 0-100
	Illuminance        int       // lux
	UV                 int       // UV index 0-11
	SolarRadiation     int       // watts per square meter
	RainAccumulation   float64   // millimeters per minute
	PrecipitationType  int       // 0=none, 1=rain, 2=hail, 3=hail+rain
	LightningStrikeAvg Distance  // average lightning strike distance in kilometers
	LightningStrikeCnt int       // lightning strike count
	BatteryVolts       float64   // sensor battery voltage
	ReportingInterval  int       // sensor reporting interval minutes
}
