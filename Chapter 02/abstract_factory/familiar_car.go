package abstract_factory

type FamiliarCar struct{}

func (l *FamiliarCar) GetDoors() int {
	return 5
}
func (l *FamiliarCar) GetWheels() int {
	return 4
}
func (l *FamiliarCar) GetSeats() int {
	return 5
}
