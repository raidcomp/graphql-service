// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type CreateUserInput struct {
	Login string `json:"login"`
	Email string `json:"email"`
}

type CreateUserPayload struct {
	User  *User                   `json:"user"`
	Error *CreateUserPayloadError `json:"error"`
}

type CreateUserPayloadError struct {
	LoginTaken   bool `json:"loginTaken"`
	LoginInvalid bool `json:"loginInvalid"`
}

type LoginUserInput struct {
	ID string `json:"id"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	ID          string    `json:"id"`
	Login       string    `json:"login"`
	Email       string    `json:"email"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}
