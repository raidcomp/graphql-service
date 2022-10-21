package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/raidcomp/graphql-service/auth"
	"github.com/raidcomp/graphql-service/graph/generated"
	"github.com/raidcomp/graphql-service/graph/model"
	"github.com/raidcomp/graphql-service/middleware"
	users_service "github.com/raidcomp/users-service/proto"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	createUserResp, err := r.UsersClient.CreateUser(ctx, &users_service.CreateUserRequest{
		Email: input.Email,
		Login: input.Login,
	})
	if err != nil {
		return nil, gqlerror.Errorf("error creating user")
	}

	token, err := auth.GenerateToken(ctx, createUserResp.User.Id)
	if err != nil {
		return nil, gqlerror.Errorf("error generating auth token for user")
	}

	return &model.CreateUserPayload{
		User: &model.User{
			ID:          createUserResp.User.Id,
			Login:       createUserResp.User.Login,
			Email:       createUserResp.User.Email,
			CreatedTime: createUserResp.User.CreatedAt.AsTime(),
			UpdatedTime: createUserResp.User.UpdatedAt.AsTime(),
		},
		Token: &token,
	}, err
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (*model.LoginUserPayload, error) {
	userID := middleware.UserID(ctx)
	if userID == "" {
		return nil, nil
	}

	_, err := r.UsersClient.CheckUserPassword(ctx, &users_service.CheckUserPasswordRequest{
		Id:       userID,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	// We're good, generate token for user
	token, err := auth.GenerateToken(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.LoginUserPayload{
		User:  nil, // TODO: get this I guess? Will need to create NewUserResolver(userID)
		Token: &token,
		Error: nil,
	}, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string, login *string) (*model.User, error) {
	var getUserResp *users_service.GetUserResponse
	var err error
	if id != nil {
		getUserResp, err = r.UsersClient.GetUser(ctx, &users_service.GetUserRequest{
			Id: *id,
		})
	} else if login != nil {
		getUserResp, err = r.UsersClient.GetUser(ctx, &users_service.GetUserRequest{
			Login: *login,
		})
	} else {
		return nil, gqlerror.Errorf("id or login are required")
	}

	if err != nil {
		return nil, gqlerror.Errorf("error requesting user")
	}

	if getUserResp.User == nil {
		return nil, nil
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
