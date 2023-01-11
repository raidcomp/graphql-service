# graphql-service

Public-Facing First-Party GraphQL service and schema for [raidcomp.io](raidcomp.io). User requests are authenticated using a [JWT scheme](https://jwt.io/).

The GraphQL service provides:
- Expressive, extensible, backwards-compatible interface for first-party clients
- Secure network traffic on the public internet
- User authentication (i.e translation and verification of the JWT token to a `user_id`)
- Request tracing

Endpoint services are responsible for:
- Resolver/mutation business logic implementation
- User authorization (i.e what a `user_id` can and cannot do)
- Rate limiting
- Business metrics 

## Technologies

- GoLang: service code
- GraphQL/gqlgen: query language to model API objects in graph-like relations, resolver, mutation, and type generator
- JWT: user authentication
- GitHub Actions: CI/CD

## Development

### `make run`

Starts the server for local development on `http://localhost:8080/gql`. Go to `http://localhost:8080/` to see the GraphiQL GUI. 

### `make generate`

Generates the GraphQL resolver and mutation endpoints based on the [schema.graphqls](/graph/schema.graphqls) schema definition.

