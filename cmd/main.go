package main

import (
	"github.com/pmorelli92/go-state-machine-two/pkg/http"
	"github.com/pmorelli92/go-state-machine-two/pkg/persistence"
	"log"
)

func main() {

	options := persistence.NewPostgresOptions()
	rp := persistence.VehicleSqlRepository{ Options: options }

	if err := http.Bootstrap(&rp); err != nil {
		log.Fatal(err)
	}
}