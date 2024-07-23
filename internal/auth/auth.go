package auth

import "go-api-starter/internal/types"

type Auth interface {
	Authenticate(types.User) (string, error)
	Authorize(string) error
}
