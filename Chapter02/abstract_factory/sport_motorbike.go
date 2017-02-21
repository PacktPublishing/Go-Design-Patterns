package abstract_factory

type SportMotorbike struct{}

func (s *SportMotorbike) GetWheels() int {
	return 2
}
func (s *SportMotorbike) GetSeats() int {
	return 1
}
func (s *SportMotorbike) GetType() int {
	return SportMotorbikeType
}
