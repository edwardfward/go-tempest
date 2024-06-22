package tempest

import (
	"encoding/json"
	"testing"
	"time"
)

func TestObservation_UnmarshalJSON(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		msg := `{"serial_number":"ST-00146014","type":"obs_st","hub_sn":"HB-00149269","obs":[[1718851666,0.00,0.00,0.00,0,15,994.17,21.15,97.04,0,0.00,0,0.000000,0,0,0,2.628,1]],"firmware_revision":176}`

		obs := Observation{}
		if err := json.Unmarshal([]byte(msg), &obs); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if obs.SerialNumber != "ST-00146014" {
			t.Errorf("unexpected serial number: %s", obs.SerialNumber)
		}

		if obs.Type != "obs_st" {
			t.Errorf("unexpected type: %s", obs.Type)
		}

		if obs.HubSerial != "HB-00149269" {
			t.Errorf("unexpected hub serial: %s", obs.HubSerial)
		}

		if !obs.Time.Equal(time.Unix(1718851666, 0)) {
			t.Errorf("expected time: %v, got: %v", time.Unix(1718851666, 0), obs.Time)
		}

		if obs.WindLull != 0.00 {
			t.Errorf("unexpected wind lull: %f", obs.WindLull)
		}

		if obs.WindAverage != 0.00 {
			t.Errorf("unexpected wind average: %f", obs.WindAverage)
		}

		if obs.WindGust != 0.00 {
			t.Errorf("unexpected wind gust: %f", obs.WindGust)
		}

		if obs.WindDirection != 0 {
			t.Errorf("unexpected wind direction: %d", obs.WindDirection)
		}

		if obs.WindSampleInterval != 15 {
			t.Errorf("unexpected wind sample interval: %d", obs.WindSampleInterval)
		}

		if obs.StationPressure != 994.17 {
			t.Errorf("unexpected station pressure: %f", obs.StationPressure)
		}

		if obs.AirTemperature != 21.15 {
			t.Errorf("unexpected air temperature: %f", obs.AirTemperature)
		}

		if obs.RelativeHumidity != 97.04 {
			t.Errorf("unexpected relative humidity: %f", obs.RelativeHumidity)
		}

		if obs.Illuminance != 0 {
			t.Errorf("unexpected illuminance: %f", obs.Illuminance)
		}

		if obs.UV != 0 {
			t.Errorf("unexpected UV: %f", obs.UV)
		}

		if obs.SolarRadiation != 0 {
			t.Errorf("unexpected solar radiation: %f", obs.SolarRadiation)
		}

		if obs.RainAccumulation != 0 {
			t.Errorf("unexpected rain accumulation: %f", obs.RainAccumulation)
		}

		if obs.PrecipitationType != 0 {
			t.Errorf("unexpected precipitation type: %d", obs.PrecipitationType)
		}

		if obs.LightningAverageDistance != 0 {
			t.Errorf("unexpected lightning average distance: %f", obs.LightningAverageDistance)
		}

		if obs.LightningStrikeCount != 0 {
			t.Errorf("unexpected lightning strike count: %d", obs.LightningStrikeCount)
		}

		if obs.Battery != 2.628 {
			t.Errorf("unexpected battery: %f", obs.Battery)
		}

		if obs.ReportingInterval != 1 {
			t.Errorf("unexpected reporting interval: %d", obs.ReportingInterval)
		}

		if obs.FirmwareRevision != 176 {
			t.Errorf("unexpected firmware revision: %d", obs.FirmwareRevision)
		}
	})
}
