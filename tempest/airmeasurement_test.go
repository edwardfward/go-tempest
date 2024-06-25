package tempest

import (
	"math"
	"testing"
)

func TestAirMeasurement_F(t *testing.T) {
	a := AirMeasurement{Temperature: TempReading{Reading: 0, Units: Celsius}}
	if a.F() != 32 {
		t.Errorf("Expected 32, got %v", a.F())
	}
}

func TestAirMeasurement_K(t *testing.T) {
	a := AirMeasurement{Temperature: TempReading{Reading: 0, Units: Celsius}}
	if a.K() != 273.15 {
		t.Errorf("Expected 273.15, got %v", a.K())
	}
}

func TestAirMeasurement_DewPointC(t *testing.T) {
	a := AirMeasurement{Temperature: TempReading{
		Reading: 0, Units: Celsius,
	}, Humidity: 100}
	if a.DewPointC() != 0 {
		t.Errorf("Expected 0, got %v", a.DewPointC())
	}
}

func TestAirMeasurement_DewPointF(t *testing.T) {
	a := AirMeasurement{Temperature: TempReading{
		Reading: 0, Units: Celsius,
	}, Humidity: 100}
	if a.DewPointF() != 32 {
		t.Errorf("Expected 32, got %v", a.DewPointF())
	}
}

func TestAirMeasurement_InHg(t *testing.T) {
	a := AirMeasurement{Pressure: PressureReading{
		pressure: 1013.25, unit: hPa,
	}}
	// Check within a reasonable margin of error.
	diff := math.Abs(a.InHg() - 29.921252)
	if diff > 0.00001 {
		t.Errorf("expected diff < 0.00001, got %v %f", diff, a.InHg())
	}
}
