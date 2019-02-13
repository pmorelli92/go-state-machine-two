package domain

import (
	"errors"
	"fmt"
	"github.com/looplab/fsm"
)

const (
	readyState = "ready"
	ridingState = "riding"
	lowBatteryState = "lowBattery"
)

const (
	startRideEvent = "startRideEvent"
	finishRideEvent = "finishRideEvent"
	lowBatteryEvent = "lowBatteryEvent"
)

type Vehicle struct {
	battery int
	fsm *fsm.FSM
}

func (v *Vehicle) Battery() int {
	return v.battery
}

func (v *Vehicle) GetCurrentState() string {
	return v.fsm.Current()
}

func (v * Vehicle) StartRide() error {
	if v.fsm.Can(startRideEvent) {
		return v.fsm.Event(startRideEvent)
	}

	errMsg := fmt.Sprintf("you cannot start a ride right now the current state is %s", v.GetCurrentState())
	return errors.New(errMsg)

}

func (v * Vehicle) FinishRide(batteryLeft int) error {
	if v.fsm.Can(finishRideEvent) {
		v.battery = batteryLeft
		if v.battery < 20 {
			return v.fsm.Event(lowBatteryEvent)
		} else {
			return v.fsm.Event(finishRideEvent)
		}
	}

	errMsg := fmt.Sprintf("you cannot finish a ride right now the current state is %s", v.GetCurrentState())
	return errors.New(errMsg)
}

func NewVehicle() *Vehicle {

	myFsm := fsm.NewFSM(
		readyState,
		fsm.Events{

			{Name: startRideEvent, Src: []string{readyState}, Dst: ridingState},
			{Name: finishRideEvent, Src: []string{ridingState}, Dst: readyState},
			{Name: lowBatteryEvent, Src: []string{ridingState}, Dst: lowBatteryState},
		},
		fsm.Callbacks{ },
	)

	return &Vehicle{battery:100, fsm:myFsm}
}
