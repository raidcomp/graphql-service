package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/raidcomp/graphql-service/graph/generated"
	"github.com/raidcomp/graphql-service/graph/model"
	users_service "github.com/raidcomp/users-service/proto"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: LoginUser - loginUser"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string, login *string) (*model.User, error) {
	idStr := *id

	getUserResp, err := r.UsersClient.GetUser(ctx, &users_service.GetUserRequest{
		Id: idStr,
	})
	if err != nil {
		return nil, gqlerror.Errorf("error requesting user")
	}

	return &model.User{
		ID:          getUserResp.User.Id,
		Login:       getUserResp.User.Login,
		Email:       getUserResp.User.Email,
		CreatedTime: getUserResp.User.CreatedAt.AsTime(),
		UpdatedTime: getUserResp.User.UpdatedAt.AsTime(),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
