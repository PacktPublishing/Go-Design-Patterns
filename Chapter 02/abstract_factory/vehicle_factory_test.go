package abstract_factory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.GetVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels and %d seats\n", motorbikeVehicle.GetWheels(), motorbikeVehicle.GetSeats())

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportBike.GetType())

	motorbikeVehicle, err = motorbikeF.GetVehicle(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.GetWheels())

	cruiseBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Cruise motorbike has type %d\n", cruiseBike.GetType())

	motorbikeVehicle, err = motorbikeF.GetVehicle(3)
	if err == nil {
		t.Fatal("Motorbike of type 3 should not be recognized")
	}
}

func TestCarFactory(t *testing.T) {
	carF, err := GetVehicleFactory(3)
	if err == nil {
		t.Fatal("Car factory with id 3 should not be recognized")
	}

	carF, err = GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.GetVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats and %d wheels\n", carVehicle.GetSeats(), carVehicle.GetWheels())

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.GetDoors())

	carVehicle, err = carF.GetVehicle(FamiliarCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats\n", carVehicle.GetWheels())

	familiarCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Familiar car has %d doors.\n", familiarCar.GetDoors())

	carVehicle, err = carF.GetVehicle(3)
	if err == nil {
		t.Fatal("Car of type 3 should not be recognized")
	}
}
