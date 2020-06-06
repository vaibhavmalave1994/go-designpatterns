package builder

import (
	"testing"
)

func TestMotorBikeBuilder(t *testing.T) {
	manufacturingDirector := ManufacturingDirector{}

	motorBikeBuilder := &MotorBikeBuilder{}
	manufacturingDirector.SetBuilder(motorBikeBuilder)
	manufacturingDirector.Construct()
	motorBike := motorBikeBuilder.GetVehicle()
	if motorBike.Wheels != 2 {
		t.Errorf("Motorbike must have 2 wheels, but it has %d\n", motorBike.Wheels)
	}

	if motorBike.Seats != 2 {
		t.Errorf("Motorbike must have 2 seats, but it has %d\n", motorBike.Seats)
	}

	if motorBike.Structure != "MotorBike" {
		t.Errorf("Structure on motorbike must be 'MotorBike', but it has %s\n", motorBike.Structure)
	}
}
