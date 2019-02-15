package main

import (
	"github.com/pmorelli92/go-state-machine-two/pkg/domain"
	"github.com/pmorelli92/go-state-machine-two/pkg/persistence"
)

func main() {

	options := persistence.NewPostgresOptions()
	rp := persistence.VehicleSqlRepository{ Options: options }

	v2, _ := rp.GetAllWithLastChangeOfStateOlderThanTwoDays()

	domain.SetVehiclesFromReadyToUnknown(v2)

	for _, v := range v2 {
		_ = rp.AddOrUpdate(*v)
	}
}