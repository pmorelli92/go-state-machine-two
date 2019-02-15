package domain

import "github.com/satori/go.uuid"

type VehicleRepository interface {
	AddOrUpdate(vehicle *Vehicle) error
	GetByID(id uuid.UUID) (*Vehicle, error)
	GetAllWhereReadyState() ([]*Vehicle, error)
	GetAllWithLastChangeOfStateOlderThanTwoDays() ([]*Vehicle, error)
}