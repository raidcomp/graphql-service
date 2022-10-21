package main

import (
	"github.com/raidcomp/graphql-service/clients"
	"github.com/raidcomp/graphql-service/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/raidcomp/graphql-service/graph"
	"github.com/raidcomp/graphql-service/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	usersClient := clients.NewUsersClient()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UsersClient: usersClient,
	}}))

	authorizationMiddleware := middleware.NewUserAuthorizationMiddleware()

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", authorizationMiddleware.AuthorizeUserRequest(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
