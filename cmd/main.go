package main

import (
	"fmt"
	"github.com/pmorelli92/go-state-machine-two/pkg/domain"
)

func main() {

	vehicle := domain.NewVehicle()

	fmt.Println(vehicle.GetCurrentState())
	err := vehicle.FinishRide(10, domain.EndUser)
	checkError(err)

	err = vehicle.StartRide(domain.EndUser)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.Ready(domain.Hunter)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.FinishRide(70, domain.EndUser)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.StartRide(domain.EndUser)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.FinishRide(19, domain.EndUser)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.Collect(domain.EndUser)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.Collect(domain.Hunter)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.Drop(domain.EndUser)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.Drop(domain.Hunter)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.Ready(domain.EndUser)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.Ready(domain.Hunter)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.SetBatteryLow(domain.Hunter)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.SetBatteryLow(domain.Admin)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	fmt.Println(vehicle.GetCurrentState())

	fmt.Println(vehicle.GetCurrentState())

	fmt.Println(vehicle.GetCurrentState())

	fmt.Println(vehicle.GetCurrentState())




}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}