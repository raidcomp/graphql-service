package graph

import users_service "github.com/raidcomp/users-service/proto"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UsersClient users_service.UsersClient
}
