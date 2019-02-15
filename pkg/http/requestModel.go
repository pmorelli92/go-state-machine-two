package http

import "github.com/pmorelli92/go-state-machine-two/pkg/domain"

type ReadyRequest struct {
	UserRole domain.UserRole
}