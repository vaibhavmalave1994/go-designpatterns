package abstractfactory

type CruiseMotorBike struct{}

func (cmb *CruiseMotorBike) GetMotorBikeType() int {
	return CruiseMotorBikeType
}

func (cmb *CruiseMotorBike) NumWheels() int {
	return 2
}

func (cmb *CruiseMotorBike) NumSeats() int {
	return 2
}
