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

type Observation struct {
	Sensor       WeatherSensor
	Hub          Hub
	Observations []Obs
}

type Obs struct {
	TimestampEpochUTC  int64
	WindLull           float64
	WindAverage        float64
	WindGust           float64
	WindDirection      float64
	WindSampleInterval int
	StationPressure    float64
	AirTemperature     float64
	RelativeHumidity   float64
	Illuminance        int
	UV                 int
	SolarRadiation     int
	RainAccumulation   float64
	PrecipitationType  int
	LightningStrikeAvg float64
	LightningStrikeCnt int
	BatteryVolts       float64
	ReportingInterval  int
}
