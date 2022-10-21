package middleware

import (
	"context"
	"github.com/raidcomp/graphql-service/auth"
	"net/http"
	"time"
)

const USER_ID_CONTEXT_KEY = "userID"
const AUTHORIZATION_HEADER_KEY = "Authorization"

type user struct {
	id string
}

type UserAuthorizationMiddleware interface {
	Handle(next http.Handler) http.Handler
}

type userAuthorizationMiddlewareImpl struct{}

func NewUserAuthorizationMiddleware() UserAuthorizationMiddleware {
	return userAuthorizationMiddlewareImpl{}
}

func (u userAuthorizationMiddlewareImpl) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get(AUTHORIZATION_HEADER_KEY)

		// Allow unauthenticated users in
		if tokenStr == "" {
			next.ServeHTTP(w, r)
			return
		}

		// Parse token
		userID, expiresAt, err := auth.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		// Validate jwt token
		// TODO: Do we validate that this is a valid userID?
		if time.Now().After(expiresAt) {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		// Put userID in context
		user := user{id: userID}
		ctx := context.WithValue(r.Context(), USER_ID_CONTEXT_KEY, &user)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func UserID(ctx context.Context) string {
	user, _ := ctx.Value(USER_ID_CONTEXT_KEY).(*user)
	return user.id
}
