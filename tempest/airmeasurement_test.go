package tempest

import (
	"math"
	"testing"
)

func TestAirMeasurement_F(t *testing.T) {
	a := AirMeasurement{Temperature: 0}
	if a.F() != 32 {
		t.Errorf("Expected 32, got %v", a.F())
	}
}

func TestAirMeasurement_K(t *testing.T) {
	a := AirMeasurement{Temperature: 0}
	if a.K() != 273.15 {
		t.Errorf("Expected 273.15, got %v", a.K())
	}
}

func TestAirMeasurement_DewPointC(t *testing.T) {
	a := AirMeasurement{Temperature: 0, Humidity: 100}
	if a.DewPointC() != 0 {
		t.Errorf("Expected 0, got %v", a.DewPointC())
	}
}

func TestAirMeasurement_DewPointF(t *testing.T) {
	a := AirMeasurement{Temperature: 0, Humidity: 100}
	if a.DewPointF() != 32 {
		t.Errorf("Expected 32, got %v", a.DewPointF())
	}
}

func TestAirMeasurement_InHg(t *testing.T) {
	a := AirMeasurement{Pressure: 1013.25}
	// Check within a reasonable margin of error.
	diff := math.Abs(a.InHg() - 29.921252)
	if diff > 0.000001 {
		t.Errorf("expected diff < 0.000001, got %v %f", diff, a.InHg())
	}
}
