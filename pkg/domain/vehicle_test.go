package domain

import (
	"testing"
	"time"
)

func TestRecreateVehicle(t *testing.T) {
	vehicle := NewVehicle()
	vehicle.fsm.SetState(ridingState)

	v2 := RecreateVehicle(vehicle.ID(), vehicle.Battery(), vehicle.LastChangeOfState(), vehicle.GetCurrentState())

	if v2.ID() != vehicle.ID() {
		t.Fail()
	}

	if v2.Battery() != vehicle.Battery() {
		t.Fail()
	}

	if v2.GetCurrentState() != vehicle.GetCurrentState() {
		t.Fail()
	}

	if v2.LastChangeOfState() != vehicle.LastChangeOfState() {
		t.Fail()
	}
}

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

	for _, userRole := range tables {

		lastState = vehicle.LastChangeOfState()
		batteryBefore = vehicle.Battery()

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

func TestCollect(t *testing.T) {
	tables := []UserRole{ Hunter, Admin }
	vehicle := NewVehicle()
	vehicle.fsm.SetState(bountyState)
	lastState := vehicle.LastChangeOfState()

	if err := vehicle.Collect(EndUser); err == nil {
		t.Fail()
	}

	if lastState != vehicle.LastChangeOfState() {
		t.Fail()
	}

	for _, userRole := range tables {
		vehicle = NewVehicle()
		vehicle.fsm.SetState(bountyState)
		lastState := vehicle.LastChangeOfState()

		if err := vehicle.Collect(userRole); err != nil {
			t.Fail()
		}

		if lastState == vehicle.LastChangeOfState() {
			t.Fail()
		}

		if vehicle.GetCurrentState() != collectedState {
			t.Fail()
		}
	}
}

func TestCollectWhenVehicleStateIsDifferentThanBounty(t *testing.T) {
	tables := []UserRole{ EndUser, Hunter }
	vehicle := NewVehicle()
	vehicle.fsm.SetState(readyState)
	lastState := vehicle.LastChangeOfState()

	if err := vehicle.Collect(Admin); err != nil {
		t.Fail() //Admin should be able to put the vehicle in any state
	}

	if lastState == vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != collectedState {
		t.Fail()
	}

	for _, userRole := range tables {

		vehicle.fsm.SetState(readyState)
		lastState = vehicle.LastChangeOfState()

		if err := vehicle.Collect(userRole); err == nil {
			t.Fail()
		}

		if lastState != vehicle.LastChangeOfState() {
			t.Fail()
		}
	}
}

func TestDrop(t *testing.T) {
	tables := []UserRole{ Hunter, Admin }
	vehicle := NewVehicle()
	vehicle.fsm.SetState(collectedState)
	lastState := vehicle.LastChangeOfState()

	if err := vehicle.Drop(EndUser); err == nil {
		t.Fail()
	}

	if lastState != vehicle.LastChangeOfState() {
		t.Fail()
	}

	for _, userRole := range tables {
		vehicle = NewVehicle()
		vehicle.fsm.SetState(collectedState)
		lastState := vehicle.LastChangeOfState()

		if err := vehicle.Drop(userRole); err != nil {
			t.Fail()
		}

		if lastState == vehicle.LastChangeOfState() {
			t.Fail()
		}

		if vehicle.GetCurrentState() != droppedState {
			t.Fail()
		}
	}
}

func TestDropWhenVehicleStateIsDifferentThanCollected(t *testing.T) {
	tables := []UserRole{ EndUser, Hunter }
	vehicle := NewVehicle()
	vehicle.fsm.SetState(readyState)
	lastState := vehicle.LastChangeOfState()

	if err := vehicle.Drop(Admin); err != nil {
		t.Fail() //Admin should be able to put the vehicle in any state
	}

	if lastState == vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != droppedState {
		t.Fail()
	}

	for _, userRole := range tables {

		vehicle.fsm.SetState(readyState)
		lastState = vehicle.LastChangeOfState()

		if err := vehicle.Drop(userRole); err == nil {
			t.Fail()
		}

		if lastState != vehicle.LastChangeOfState() {
			t.Fail()
		}
	}
}

func TestReady(t *testing.T) {
	tables := []UserRole{ Hunter, Admin }
	vehicle := NewVehicle()
	vehicle.fsm.SetState(droppedState)
	lastState := vehicle.LastChangeOfState()

	if err := vehicle.Ready(EndUser); err == nil {
		t.Fail()
	}

	if lastState != vehicle.LastChangeOfState() {
		t.Fail()
	}

	for _, userRole := range tables {
		vehicle = NewVehicle()
		vehicle.fsm.SetState(droppedState)
		lastState := vehicle.LastChangeOfState()

		if err := vehicle.Ready(userRole); err != nil {
			t.Fail()
		}

		if lastState == vehicle.LastChangeOfState() {
			t.Fail()
		}

		if vehicle.GetCurrentState() != readyState {
			t.Fail()
		}
	}
}

func TestReadyWhenVehicleStateIsDifferentThanDropped(t *testing.T) {
	tables := []UserRole{ EndUser, Hunter }
	vehicle := NewVehicle()
	vehicle.fsm.SetState(readyState)
	lastState := vehicle.LastChangeOfState()

	if err := vehicle.Ready(Admin); err != nil {
		t.Fail() //Admin should be able to put the vehicle in any state
	}

	if lastState == vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != readyState {
		t.Fail()
	}

	for _, userRole := range tables {

		vehicle.fsm.SetState(readyState)
		lastState = vehicle.LastChangeOfState()

		if err := vehicle.Ready(userRole); err == nil {
			t.Fail()
		}

		if lastState != vehicle.LastChangeOfState() {
			t.Fail()
		}
	}
}

func TestBatteryLow(t *testing.T) {
	tables := []UserRole{ Hunter, EndUser }
	vehicle := NewVehicle()
	lastState := vehicle.LastChangeOfState()

	if err := vehicle.SetBatteryLow(Admin); err != nil {
		t.Fail() //Admin should be able to put the vehicle in any state
	}

	if lastState == vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != batteryLowState {
		t.Fail()
	}

	for _, userRole := range tables {
		vehicle = NewVehicle()
		lastState := vehicle.LastChangeOfState()

		if err := vehicle.SetBatteryLow(userRole); err == nil {
			t.Fail()
		}

		if lastState != vehicle.LastChangeOfState() {
			t.Fail()
		}

		if vehicle.GetCurrentState() != readyState {
			t.Fail()
		}
	}
}

func TestSetBounty(t *testing.T) {
	tables := []UserRole{ Hunter, EndUser }
	vehicle := NewVehicle()
	lastState := vehicle.LastChangeOfState()

	if err := vehicle.SetBounty(Admin); err != nil {
		t.Fail() //Admin should be able to put the vehicle in any state
	}

	if lastState == vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != bountyState {
		t.Fail()
	}

	for _, userRole := range tables {
		vehicle = NewVehicle()
		lastState := vehicle.LastChangeOfState()

		if err := vehicle.SetBounty(userRole); err == nil {
			t.Fail()
		}

		if lastState != vehicle.LastChangeOfState() {
			t.Fail()
		}

		if vehicle.GetCurrentState() != readyState {
			t.Fail()
		}
	}
}

func TestUnknown(t *testing.T) {
	tables := []UserRole{ Hunter, EndUser }
	vehicle := NewVehicle()
	lastState := vehicle.LastChangeOfState()

	if err := vehicle.Unknown(Admin); err != nil {
		t.Fail() //Admin should be able to put the vehicle in any state
	}

	if lastState == vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != unknownState {
		t.Fail()
	}

	for _, userRole := range tables {
		vehicle = NewVehicle()
		lastState := vehicle.LastChangeOfState()

		if err := vehicle.Unknown(userRole); err == nil {
			t.Fail()
		}

		if lastState != vehicle.LastChangeOfState() {
			t.Fail()
		}

		if vehicle.GetCurrentState() != readyState {
			t.Fail()
		}
	}
}

func TestSetVehiclesFromReadyToBounty(t *testing.T) {
	vehicle := NewVehicle()
	lastState := vehicle.LastChangeOfState()

	if err := SetVehiclesFromReadyToBounty([]*Vehicle { vehicle }); err != nil {
		t.Fail()
	}

	if lastState == vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != bountyState {
		t.Fail()
	}
}

func TestSetVehiclesFromReadyToBountyWhenVehicleStateIsDifferentThanReady(t *testing.T) {
	vehicle := NewVehicle()
	vehicle.fsm.SetState(ridingState)
	lastState := vehicle.LastChangeOfState()

	if err := SetVehiclesFromReadyToBounty([]*Vehicle { vehicle }); err == nil {
		t.Fail()
	}

	if lastState != vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != ridingState {
		t.Fail()
	}
}

func TestSetVehiclesFromReadyToUnknown(t *testing.T) {
	vehicle := NewVehicle()

	// We need to make it last change of state 48 before now
	lastState := time.Now().AddDate(0, 0, -2)
	vehicle.lastChangeOfState = lastState

	if err := SetVehiclesFromReadyToUnknown([]*Vehicle { vehicle }); err != nil {
		t.Fail()
	}

	if lastState == vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != unknownState {
		t.Fail()
	}
}

func TestSetVehiclesFromReadyToUnknownWhenVehicleLastChangeOfStateHappenedBefore48Hours(t *testing.T) {
	vehicle := NewVehicle()
	lastState := vehicle.LastChangeOfState()

	if err := SetVehiclesFromReadyToUnknown([]*Vehicle { vehicle }); err == nil {
		t.Fail()
	}

	if lastState != vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != readyState {
		t.Fail()
	}
}

func TestSetVehiclesFromReadyToUnknownWhenVehicleStateIsDifferentThanReady(t *testing.T) {
	vehicle := NewVehicle()
	vehicle.fsm.SetState(ridingState)

	// We need to make it last change of state 48 before now
	lastState := time.Now().AddDate(0, 0, -2)
	vehicle.lastChangeOfState = lastState

	if err := SetVehiclesFromReadyToUnknown([]*Vehicle { vehicle }); err == nil {
		t.Fail()
	}

	if lastState != vehicle.LastChangeOfState() {
		t.Fail()
	}

	if vehicle.GetCurrentState() != ridingState {
		t.Fail()
	}
}