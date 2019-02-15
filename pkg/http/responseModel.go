package http

import (
	"github.com/pmorelli92/go-state-machine-two/pkg/domain"
	"github.com/satori/go.uuid"
	"time"
)

type ResourceResponse struct {
	Id uuid.UUID
}

type ErrorResponse struct {
	Message string
}

type VehicleResponse struct {
	Id uuid.UUID
	Battery int
	LastChangeOfState time.Time
	CurrentState string
}

func ToResponseModel(vehicle *domain.Vehicle) VehicleResponse {
	return VehicleResponse{
		Id:vehicle.Id(),
		Battery:vehicle.Battery(),
		LastChangeOfState:vehicle.LastChangeOfState(),
		CurrentState:vehicle.GetCurrentState()}
}