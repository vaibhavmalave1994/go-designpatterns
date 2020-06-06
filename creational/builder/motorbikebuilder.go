package builder

//Motorbike builder starts here
type MotorBikeBuilder struct {
	vehicle VehicleProduct
}

func (m *MotorBikeBuilder) SetWheels() BuildProcess {
	m.vehicle.Wheels = 2
	return m
}

func (m *MotorBikeBuilder) SetSeats() BuildProcess {
	m.vehicle.Seats = 2
	return m
}

func (m *MotorBikeBuilder) SetStructure() BuildProcess {
	m.vehicle.Structure = "MotorBike"
	return m
}

func (m *MotorBikeBuilder) GetVehicle() VehicleProduct {
	return m.vehicle
}

//Motorbike builder ends here
