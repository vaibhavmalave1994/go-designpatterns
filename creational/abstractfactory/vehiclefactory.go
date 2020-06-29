package abstractfactory

import (
	"errors"
	"fmt"
)

const (
	CarFactoryType       = 1
	MotorBikeFactoryType = 2
)

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

func BuildFactory(vtype int) (VehicleFactory, error) {
	switch vtype {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorBikeFactoryType:
		return new(MotorBikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory of type %d is not recognised.\n", vtype))

	}
}
