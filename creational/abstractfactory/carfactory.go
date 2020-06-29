package abstractfactory

import (
	"errors"
	"fmt"
)

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (cf *CarFactory) NewVehicle(vtype int) (Vehicle, error) {
	switch vtype {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d is not recognised.\n", vtype))
	}
	return nil, nil
}
