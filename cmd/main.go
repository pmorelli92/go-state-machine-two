package main

import (
	"fmt"
	"github.com/pmorelli92/go-state-machine-two/pkg/domain"
)

func main() {

	vehicle := domain.NewVehicle()

	fmt.Println(vehicle.GetCurrentState())
	err := vehicle.FinishRide(10)
	checkError(err)

	err = vehicle.StartRide()
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.FinishRide(70)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.StartRide()
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

	err = vehicle.FinishRide(19)
	checkError(err)
	fmt.Println(vehicle.GetCurrentState())

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}