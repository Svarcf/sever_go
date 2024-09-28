package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/Svarcf/sever_go/internal/services"
)

type Resolver struct {
	UserService *services.UserService
}
