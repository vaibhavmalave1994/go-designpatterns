package abstractfactory

import (
	"errors"
	"fmt"
)

const (
	SportMotorBikeType  = 1
	CruiseMotorBikeType = 2
)

type MotorBikeFactory struct{}

func (mbf *MotorBikeFactory) NewVehicle(vtype int) (Vehicle, error) {
	switch vtype {
	case SportMotorBikeType:
		return new(SportMotorBike), nil
	case CruiseMotorBikeType:
		return new(CruiseMotorBike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d is not recongnised\n", vtype))
	}
}
