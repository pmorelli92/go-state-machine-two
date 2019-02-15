package http

import "github.com/pmorelli92/go-state-machine-two/pkg/domain"

type BaseRequest struct {
	UserRole    domain.UserRole
	BatteryLeft int
}
