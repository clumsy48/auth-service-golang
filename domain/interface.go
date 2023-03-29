package domain

import (
	"context"
)

type AuthService interface {
	Authenticate(context.Context, AuthParams) (AuthTokenID, error)
	Validate(context.Context, AuthTokenID) (string, error)
	Expire(context.Context, string) error
}

type OnboardingService interface {
	SignUp(context.Context, User) error
	Delete(context.Context, AuthTokenID) error
}

type Storage interface {
	Save(context.Context, Credential) error
	SaveUser(context.Context, User) error
	FindAll(context.Context) ([]Credential, error)
	FindByID(context.Context, CredentialID) (Credential, error)
	FindByEmail(context.Context, string) (User, error)
	FindByAuthTokenID(context.Context, AuthTokenID) (Credential, error)
	DeleteByID(context.Context, CredentialID) error
}

type Hasher interface {
	Hash(string) (string, error)
	Compare(password, hash string) bool
}
