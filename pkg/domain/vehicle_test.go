package domain

import "testing"

func TestStartRide(t *testing.T) {
	tables := []UserRole{ EndUser, Hunter, Admin }
	for _, userRole := range tables {
		vehicle := NewVehicle()
		lastState := vehicle.LastChangeOfState()

		if err := vehicle.StartRide(userRole); err != nil {
			t.Fail()
		}

		if lastState == vehicle.LastChangeOfState() {
			t.Fail()
		}

		if vehicle.GetCurrentState() != ridingState {
			t.Fail()
		}
	}
}

func TestStartRideWhenVehicleStateIsDifferentThanReady(t *testing.T) {
	tables := []UserRole{ EndUser, Hunter }
	vehicle := NewVehicle()
	vehicle.fsm.SetState(ridingState)
	lastState := vehicle.LastChangeOfState()

	if err := vehicle.StartRide(Admin); err != nil {
		t.Fail() //Admin should be able to put the vehicle in any state
	}

	if lastState == vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != ridingState {
		t.Fail()
	}

	lastState = vehicle.LastChangeOfState()

	for _, userRole := range tables {

		if err := vehicle.StartRide(userRole); err == nil {
			t.Fail()
		}

		if lastState != vehicle.LastChangeOfState() {
			t.Fail()
		}
	}
}

func TestFinishRide(t *testing.T) {
	tables := []UserRole{ EndUser, Hunter, Admin }
	for _, userRole := range tables {
		vehicle := NewVehicle()
		vehicle.fsm.SetState(ridingState)
		lastState := vehicle.LastChangeOfState()
		batteryBefore := vehicle.Battery()

		if err := vehicle.FinishRide(70, userRole); err != nil {
			t.Fail()
		}

		if lastState == vehicle.LastChangeOfState() {
			t.Fail()
		}

		if batteryBefore == vehicle.Battery() {
			t.Fail()
		}

		if vehicle.GetCurrentState() != readyState {
			t.Fail()
		}
	}
}

func TestFinishRideWithLowBattery(t *testing.T) {
	tables := []UserRole{ EndUser, Hunter, Admin }
	for _, userRole := range tables {
		vehicle := NewVehicle()
		vehicle.fsm.SetState(ridingState)
		lastState := vehicle.LastChangeOfState()
		batteryBefore := vehicle.Battery()

		if err := vehicle.FinishRide(19, userRole); err != nil {
			t.Fail()
		}

		if lastState == vehicle.LastChangeOfState() {
			t.Fail()
		}

		if batteryBefore == vehicle.Battery() {
			t.Fail()
		}

		if vehicle.GetCurrentState() != bountyState {
			t.Fail()
		}
	}
}

func TestFinishRideWhenVehicleStateIsDifferentThanRiding(t *testing.T) {
	tables := []UserRole{ EndUser, Hunter }
	vehicle := NewVehicle()
	vehicle.fsm.SetState(readyState)
	lastState := vehicle.LastChangeOfState()
	batteryBefore := vehicle.Battery()

	if err := vehicle.FinishRide(70, Admin); err != nil {
		t.Fail() //Admin should be able to put the vehicle in any state
	}

	if lastState == vehicle.LastChangeOfState() {
		t.Fail()
	}

	if batteryBefore == vehicle.Battery() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != readyState {
		t.Fail()
	}

	lastState = vehicle.LastChangeOfState()
	batteryBefore = vehicle.Battery()

	for _, userRole := range tables {

		if err := vehicle.FinishRide(70, userRole); err == nil {
			t.Fail()
		}

		if lastState != vehicle.LastChangeOfState() {
			t.Fail()
		}

		if batteryBefore != vehicle.Battery() {
			t.Fail()
		}
	}
}