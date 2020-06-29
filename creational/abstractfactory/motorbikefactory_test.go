package abstractfactory

import (
	"testing"
)

func TestMotorBikeFactory(t *testing.T) {
	motorBikeFactory, err := BuildFactory(MotorBikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorBikeVehicle, err := motorBikeFactory.NewVehicle(SportMotorBikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Sport bike has %d of doors and %d of seats", motorBikeVehicle.NumWheels(), motorBikeVehicle.NumSeats())

	sportBike, ok := motorBikeVehicle.(MotorBike)
	if !ok {
		t.Errorf("Structure assertion has failed.")
	}

	t.Logf("Sport bike has %d type", sportBike.GetMotorBikeType())

	motorBikeVehicle, err = motorBikeFactory.NewVehicle(CruiseMotorBikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Family car has %d of doors and %d of seats", motorBikeVehicle.NumWheels(), motorBikeVehicle.NumSeats())

	cruiseCar, ok := motorBikeVehicle.(MotorBike)
	if !ok {
		t.Errorf("Structure assertion has failed.")
	}

	t.Logf("Family car has %d doors", cruiseCar.GetMotorBikeType())

	motorBikeVehicle, err = motorBikeFactory.NewVehicle(3)
	if err == nil {
		t.Fatal("Car of type 3 should not be recognized")
	}

}
