package clients

import users_service "github.com/raidcomp/users-service/proto"

type Clients struct {
	UsersClient users_service.UsersClient
}
