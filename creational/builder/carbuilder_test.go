package builder

import (
	"testing"
)

func TestCarBuilder(t *testing.T) {
	manufacturingDirector := ManufacturingDirector{}

	carbuilder := &CarBuilder{}
	manufacturingDirector.SetBuilder(carbuilder)
	manufacturingDirector.Construct()
	car := carbuilder.GetVehicle()
	if car.Wheels != 4 {
		t.Errorf("Car must have 4 wheels, but it has %d\n", car.Wheels)
	}

	if car.Seats != 4 {
		t.Errorf("Car must have 4 seats, but it has %d\n", car.Seats)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on car must be 'Car', but it has %s\n", car.Structure)
	}
}
