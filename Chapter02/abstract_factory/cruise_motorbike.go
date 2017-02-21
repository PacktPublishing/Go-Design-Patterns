package abstract_factory

type CruiseMotorbike struct{}

func (c *CruiseMotorbike) GetWheels() int {
	return 2
}
func (c *CruiseMotorbike) GetSeats() int {
	return 2
}
func (c *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}
