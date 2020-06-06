package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type ManufacturingDirector struct {
	Builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	//implementation goes here
	f.Builder.SetWheels().SetSeats().SetStructure()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	//implementation goes here
	f.Builder = b
}
