package tempest

import "testing"

func FuzzType_Valid(f *testing.F) {
	f.Add("MessageTypeRainStartEvent")
	f.Add("MessageTypeLightningStrike")
	f.Add("MessageTypeRapidWind")
	f.Add("MessageTypeObservation")
	f.Add("MessageTypeAirObservation")
	f.Add("MessageTypeSkyObservation")
	f.Add("MessageTypeDeviceStatus")
	f.Add("MessageTypeHubStatus")
	f.Fuzz(func(t *testing.T, s string) {
		if _, ok := validTypes[Type(s)]; ok {
			if !Type(s).Valid() {
				t.Fatalf("expected %s to be valid", s)
			}
		} else {
			if Type(s).Valid() {
				t.Fatalf("expected %s to be invalid", s)
			}
		}
	})
}
