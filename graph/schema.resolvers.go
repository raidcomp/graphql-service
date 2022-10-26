package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/raidcomp/graphql-service/auth"
	"github.com/raidcomp/graphql-service/graph/generated"
	"github.com/raidcomp/graphql-service/graph/model"
	"github.com/raidcomp/graphql-service/graph/resolvers"
	users_service "github.com/raidcomp/users-service/proto"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	createUserResp, err := r.Clients.UsersClient.CreateUser(ctx, &users_service.CreateUserRequest{
		Email:    input.Email,
		Login:    input.Login,
		Password: input.Password,
	})
	if err != nil {
		return nil, gqlerror.Errorf("error creating user")
	}

	token, err := auth.GenerateToken(ctx, createUserResp.User.Id)
	// TODO: Handle error better

	return &model.CreateUserPayload{
		User:  resolvers.NewUserResolver(createUserResp.User),
		Token: &token,
	}, err
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (*model.LoginUserPayload, error) {
	_, err := r.Clients.UsersClient.CheckUserPassword(ctx, &users_service.CheckUserPasswordRequest{
		Login:    input.Login,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	getUserResponse, err := r.Clients.UsersClient.GetUser(ctx, &users_service.GetUserRequest{
		Login: input.Login,
	})
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(ctx, getUserResponse.User.Id)
	// TODO: Handle error better

	return &model.LoginUserPayload{
		User:  resolvers.NewUserResolver(getUserResponse.User),
		Token: &token,
		Error: nil,
	}, err
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (*model.RefreshTokenPayload, error) {
	userID, _, err := auth.ParseToken(input.Token)
	if err != nil {
		return nil, gqlerror.Errorf("error parsing token")
	}

	if userID == "" {
		return nil, gqlerror.Errorf("no token provided to refresh")
	}

	newToken, err := auth.GenerateToken(ctx, userID)
	if err != nil {
		return nil, gqlerror.Errorf("error generating token")
	}

	user, err := resolvers.NewUserResolverByID(ctx, r.Clients, userID)
	if err != nil {
		return nil, err
	}

	return &model.RefreshTokenPayload{
		User:  user,
		Token: &newToken,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string, login *string) (*model.User, error) {
	var user *model.User
	var err error
	if id != nil {
		user, err = resolvers.NewUserResolverByID(ctx, r.Clients, *id)
	} else if login != nil {
		user, err = resolvers.NewUserResolverByLogin(ctx, r.Clients, *login)
	} else {
		return nil, gqlerror.Errorf("id or login are required")
	}

	if err != nil {
		return nil, gqlerror.Errorf("error requesting user")
	}

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
