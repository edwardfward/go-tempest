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
	Sensor       WeatherSensor
	Hub          Hub
	Observations []WeatherObservation
}

// WeatherObservation message type.
type WeatherObservation struct {
	EpochSecondsUTC    int64       // Epoch seconds UTC
	WindLull           Speed       // m/s
	WindAverage        Speed       // m/s
	WindGust           Speed       // m/s
	WindDirection      float64     // degrees
	WindSampleInterval int         // seconds
	StationPressure    Pressure    // millibars
	AirTemperature     TempReading // degrees Celsius
	RelativeHumidity   float64     // percentage 0-100
	Illuminance        int         // lux
	UV                 int         // UV index 0-11
	SolarRadiation     int         // W/m2
	RainAccumulation   float64
	PrecipitationType  int
	LightningStrikeAvg float64
	LightningStrikeCnt int
	BatteryVolts       float64
	ReportingInterval  int
}
