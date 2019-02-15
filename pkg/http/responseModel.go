package http

import (
	"github.com/pmorelli92/go-state-machine-two/pkg/domain"
	"github.com/satori/go.uuid"
	"time"
)

type ResourceResponse struct {
	ID uuid.UUID `json:"Id"`
}

type ErrorResponse struct {
	Message string
}

func ToErrorResponseModel(errors []error) []ErrorResponse {
	var rsp []ErrorResponse
	for _, e := range errors {
		rsp = append(rsp, ErrorResponse{Message:e.Error()})
	}
	return rsp
}

type VehicleResponse struct {
	ID                uuid.UUID `json:"Id"`
	Battery           int
	LastChangeOfState time.Time
	CurrentState      string
}

func ToResponseModel(vehicle *domain.Vehicle) VehicleResponse {
	return VehicleResponse{
		ID:                vehicle.ID(),
		Battery:           vehicle.Battery(),
		LastChangeOfState: vehicle.LastChangeOfState(),
		CurrentState:      vehicle.GetCurrentState()}
}