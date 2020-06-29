package abstractfactory

type SportMotorBike struct{}

func (smb *SportMotorBike) GetMotorBikeType() int {
	return SportMotorBikeType
}

func (smb *SportMotorBike) NumWheels() int {
	return 2
}

func (smb *SportMotorBike) NumSeats() int {
	return 1
}
