package abstract_factory

type LuxuryCar struct{}

func (l *LuxuryCar) GetDoors() int {
	return 4
}
func (l *LuxuryCar) GetWheels() int {
	return 4
}
func (l *LuxuryCar) GetSeats() int {
	return 5
}
