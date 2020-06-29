package abstractfactory

import (
	"testing"
)

func TestCarFactory(t *testing.T) {
	carFactory, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carFactory.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Luxury car has %d of doors and %d of seats", carVehicle.NumWheels(), carVehicle.NumSeats())

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Errorf("Structure assertion has failed.")
	}

	t.Logf("Luxury car has %d doors", luxuryCar.NumDoors())

	carVehicle, err = carFactory.NewVehicle(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Family car has %d of doors and %d of seats", carVehicle.NumWheels(), carVehicle.NumSeats())

	familyCar, ok := carVehicle.(Car)
	if !ok {
		t.Errorf("Structure assertion has failed.")
	}

	t.Logf("Family car has %d doors", familyCar.NumDoors())

	carVehicle, err = carFactory.NewVehicle(3)
	if err == nil {
		t.Fatal("Car of type 3 should not be recognized")
	}

}
