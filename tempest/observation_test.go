package tempest

import (
	"testing"
	"time"
)

func TestObservation_Read(t *testing.T) {
	rawMessage := `{"serial_number":"ST-00146014","type":"obs_st","hub_sn":"HB-00149269","obs":[[1719767641,0.31,1.71,3.15,358,3,995.90,18.68,57.51,159176,12.46,1326,0.000000,0,0,0,2.755,1]],"firmware_revision":176}`
	obs := Observation{}
	err := obs.Read([]byte(rawMessage))
	if err != nil {
		t.Errorf("error reading observation: %v", err)
	}

	if obs.SensorSerial != "ST-00146014" {
		t.Errorf("unexpected serial number: %s", obs.SensorSerial)
	}

	if obs.HubSerial != "HB-00149269" {
		t.Errorf("unexpected hub serial number: %s", obs.HubSerial)
	}

	if len(obs.Observations) != 1 {
		t.Fatalf("unexpected number of observations: %d", len(obs.Observations))
	}

	obs1 := obs.Observations[0]
	if obs1.EpochSecondsUTC.Equal(time.Unix(1719767641, 0)) != true {
		t.Errorf("unexpected epoch seconds UTC: %d", obs1.EpochSecondsUTC.Unix())
	}

	if obs1.WindLull.MetersPerSecond() != 0.31 {
		t.Errorf("unexpected wind lull: %f", obs1.WindLull.MetersPerSecond())
	}

	if obs1.WindAverage.MetersPerSecond() != 1.71 {
		t.Errorf("unexpected wind average: %f", obs1.WindAverage.MetersPerSecond())
	}

	if obs1.WindGust.MetersPerSecond() != 3.15 {
		t.Errorf("unexpected wind gust: %f", obs1.WindGust.MetersPerSecond())
	}

	if obs1.WindDirection.Degrees() != 358 {
		t.Errorf("unexpected wind direction: %f", obs1.WindDirection.Degrees())
	}

	if obs1.WindDirection.Cardinal() != "N" {
		t.Errorf("unexpected wind direction: %s", obs1.WindDirection.Cardinal())
	}

	if obs1.WindSampleInterval != 3 {
		t.Errorf("unexpected wind sample interval: %d", obs1.WindSampleInterval)
	}

	if obs1.StationPressure.Millibar() != 995.90 {
		t.Errorf("unexpected station pressure: %f", obs1.StationPressure.Millibar())
	}

	if obs1.AirTemperature.C() != 18.68 {
		t.Errorf("unexpected air temperature: %f", obs1.AirTemperature.C())
	}

	if obs1.RelativeHumidity != 57.51 {
		t.Errorf("unexpected relative humidity: %f", obs1.RelativeHumidity)
	}

	if obs1.Illuminance != 159176 {
		t.Errorf("unexpected illuminance: %d", obs1.Illuminance)
	}

	if obs1.UV != 12.46 {
		t.Errorf("unexpected UV index: %f", obs1.UV)
	}

	if obs1.SolarRadiation != 1326 {
		t.Errorf("unexpected solar radiation: %d", obs1.SolarRadiation)
	}

	if obs1.RainAccumulation != 0.0 {
		t.Errorf("unexpected rain accumulation: %f", obs1.RainAccumulation)
	}

	if obs1.PrecipitationType != 0 {
		t.Errorf("unexpected precipitation type: %d", obs1.PrecipitationType)
	}

	if obs1.LightningStrikeAvg.Kilometers() != 0.0 {
		t.Errorf("unexpected lightning strike average distance: %f", obs1.LightningStrikeAvg.Kilometers())
	}

	if obs1.LightningStrikeCnt != 0 {
		t.Errorf("unexpected lightning strike count: %d", obs1.LightningStrikeCnt)
	}

	if obs1.BatteryVolts != 2.755 {
		t.Errorf("unexpected battery volts: %f", obs1.BatteryVolts)
	}

	if obs1.ReportingInterval != 1 {
		t.Errorf("unexpected reporting interval: %d", obs1.ReportingInterval)
	}
}
