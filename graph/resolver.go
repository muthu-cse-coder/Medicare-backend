package graph

import (
	"medicare-backend/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

// Resolver is the root resolver
type Resolver struct {
	AuthService *service.AuthService
}

// NewResolver creates a new resolver with all services
func NewResolver() *Resolver {
	return &Resolver{
		AuthService: service.NewAuthService(),
	}
}
