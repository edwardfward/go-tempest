package tempest

import (
	"encoding/json"
	"fmt"
	"time"
)

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

// Read parses the observation message.
func (o *Observation) Read(message []byte) error {
	var rawMessage map[string]interface{}
	if err := json.Unmarshal(message, &rawMessage); err != nil {
		return err
	}

	// Clear the observations.
	o.Observations = make([]WeatherObservation, 0)

	// Record the sensor and hub serial numbers.
	// These are used to identify the sensor and hub that sent
	// the observation.
	o.SensorSerial = rawMessage["serial_number"].(string)
	o.HubSerial = rawMessage["hub_sn"].(string)

	// A message may have one or more observations.
	// The majority of the time there is only one observation.
	obs := rawMessage["obs"].([]interface{})
	for _, ob := range obs {
		observation := ob.([]interface{})
		weatherObs := WeatherObservation{}

		// Epoch seconds UTC the observation was taken.
		epochUtc, ok := observation[obsIndexTimestampEpochUTC].(float64)
		if !ok {
			return fmt.Errorf("unable to read epoch seconds utc")
		}
		weatherObs.EpochSecondsUTC = time.Unix(int64(epochUtc), 0)

		// Wind lull.
		windLull, ok := observation[obsIndexWindLull].(float64)
		if !ok {
			return fmt.Errorf("unable to read wind lull")
		}
		weatherObs.WindLull = NewSpeed(windLull, MetersPerSecond)

		// Wind average.
		windAverage, ok := observation[obsIndexWindAverage].(float64)
		if !ok {
			return fmt.Errorf("unable to read wind average")
		}
		weatherObs.WindAverage = NewSpeed(windAverage, MetersPerSecond)

		// Wind gust.
		windGust, ok := observation[obsIndexWindGust].(float64)
		if !ok {
			return fmt.Errorf("unable to read wind gust")
		}
		weatherObs.WindGust = NewSpeed(windGust, MetersPerSecond)

		// Wind direction.
		windDirection, ok := observation[obsIndexWindDirection].(float64)
		if !ok {
			return fmt.Errorf("unable to read wind direction")
		}
		weatherObs.WindDirection = NewDirection(windDirection, Degrees)

		// Wind sample interval.
		windSampleInterval, ok := observation[obsIndexWindSampleInterval].(float64)
		if !ok {
			return fmt.Errorf("unable to read wind sample interval")
		}
		weatherObs.WindSampleInterval = int(windSampleInterval)

		// Station pressure.
		stationPressure, ok := observation[obsIndexStationPressure].(float64)
		if !ok {
			return fmt.Errorf("unable to read station pressure")
		}
		weatherObs.StationPressure = NewPressure(stationPressure, Millibar)

		// Air temperature.
		airTemperature, ok := observation[obsIndexAirTemperature].(float64)
		if !ok {
			return fmt.Errorf("unable to read air temperature")
		}
		weatherObs.AirTemperature = NewTemp(airTemperature, Celsius)

		// Relative humidity.
		relativeHumidity, ok := observation[obsIndexRelativeHumidity].(float64)
		if !ok {
			return fmt.Errorf("unable to read relative humidity")
		}
		weatherObs.RelativeHumidity = relativeHumidity

		// Illuminance.
		illuminance, ok := observation[obsIndexIlluminance].(float64)
		if !ok {
			return fmt.Errorf("unable to read illuminance")
		}
		weatherObs.Illuminance = int(illuminance)

		// UV index.
		uv, ok := observation[obsIndexUV].(float64)
		if !ok {
			return fmt.Errorf("unable to read uv index")
		}
		weatherObs.UV = uv

		// Solar radiation.
		solarRadiation, ok := observation[obsIndexSolarRadiation].(float64)
		if !ok {
			return fmt.Errorf("unable to read solar radiation")
		}
		weatherObs.SolarRadiation = int(solarRadiation)

		// Rain accumulation.
		rainAccumulation, ok := observation[obsIndexRainAccumulationPastMinute].(float64)
		if !ok {
			return fmt.Errorf("unable to read rain accumulation")
		}
		weatherObs.RainAccumulation = rainAccumulation

		// Precipitation type.
		precipitationType, ok := observation[obsIndexPrecipitationType].(float64)
		if !ok {
			return fmt.Errorf("unable to read precipitation type")
		}
		weatherObs.PrecipitationType = int(precipitationType)

		// Lightning strike average distance.
		lightningStrikeAvg, ok := observation[obsIndexLightningStrikeAverageDistance].(float64)
		if !ok {
			return fmt.Errorf("unable to read lightning strike average distance")
		}
		weatherObs.LightningStrikeAvg = NewDistance(lightningStrikeAvg, Kilometers)

		// Lightning strike count.
		lightningStrikeCnt, ok := observation[obsIndexLightingStrikeCount].(float64)
		if !ok {
			return fmt.Errorf("unable to read lightning strike count")
		}
		weatherObs.LightningStrikeCnt = int(lightningStrikeCnt)

		// Battery volts.
		batteryVolts, ok := observation[obsIndexBatteryVolts].(float64)
		if !ok {
			return fmt.Errorf("unable to read battery volts")
		}
		weatherObs.BatteryVolts = batteryVolts

		// Reporting interval.
		reportingInterval, ok := observation[obsIndexReportingInterval].(float64)
		if !ok {
			return fmt.Errorf("unable to read reporting interval")
		}
		weatherObs.ReportingInterval = int(reportingInterval)

		// Append the observation to the list.
		o.Observations = append(o.Observations, weatherObs)
	}

	return nil
}

// WeatherObservation message type.
type WeatherObservation struct {
	EpochSecondsUTC    time.Time // Epoch seconds UTC
	WindLull           Speed     // meters per second
	WindAverage        Speed     // meters per second
	WindGust           Speed     // meters per second
	WindDirection      Direction // degrees
	WindSampleInterval int       // seconds
	StationPressure    Pressure  // millibars
	AirTemperature     Temp      // degrees Celsius
	RelativeHumidity   float64   // percentage 0-100
	Illuminance        int       // lux
	UV                 float64   // UV index 0-11
	SolarRadiation     int       // watts per square meter
	RainAccumulation   float64   // millimeters per minute
	PrecipitationType  int       // 0=none, 1=rain, 2=hail, 3=hail+rain
	LightningStrikeAvg Distance  // average lightning strike distance in kilometers
	LightningStrikeCnt int       // lightning strike count
	BatteryVolts       float64   // sensor battery voltage
	ReportingInterval  int       // sensor reporting interval minutes
}
