package resolvers

import (
	"context"
	"github.com/raidcomp/graphql-service/clients"
	"github.com/raidcomp/graphql-service/graph/model"
	users_service "github.com/raidcomp/users-service/proto"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func NewUserResolver(user *users_service.User) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		ID:          user.Id,
		Login:       user.Login,
		Email:       user.Email,
		CreatedTime: user.CreatedAt.AsTime(),
		UpdatedTime: user.UpdatedAt.AsTime(),
	}
}
func NewUserResolverByID(ctx context.Context, clients clients.Clients, id string) (*model.User, error) {
	getUserResp, err := clients.UsersClient.GetUser(ctx, &users_service.GetUserRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return NewUserResolver(getUserResp.User), nil
}
func NewUserResolverByLogin(ctx context.Context, clients clients.Clients, login string) (*model.User, error) {
	getUserResp, err := clients.UsersClient.GetUser(ctx, &users_service.GetUserRequest{
		Login: login,
	})

	if err != nil {
		return nil, gqlerror.Errorf("error requesting user")
	}

	return NewUserResolver(getUserResp.User), nil
}
