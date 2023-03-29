package domain

import "C"
import (
	"context"
	"fmt"
)

type AuthServiceImpl struct {
	hasher  Hasher
	storage Storage
}

func NewAuthService(s Storage, h Hasher) *AuthServiceImpl {
	return &AuthServiceImpl{storage: s, hasher: h}
}

func (as AuthServiceImpl) Authenticate(ctx context.Context, ap AuthParams) (AuthTokenID, error) {
	user, err := as.storage.FindByEmail(ctx, ap.Email)

	if err != nil {
		return "", fmt.Errorf("credentials not found")
	}

	if !as.hasher.Compare(ap.Password, user.Password) {
		return "", fmt.Errorf("credentials do not match")
	}

	cred := Credential{
		User: fmt.Sprintf("%v", user.ID),
	}
	if err := cred.GenerateAuthToken(); err != nil {
		return "", fmt.Errorf("generate auth token failed")
	}

	if err = as.storage.Save(ctx, cred); err != nil {
		return "", fmt.Errorf("save credential failed")
	}

	return cred.AuthToken.ID, nil
}

func (as AuthServiceImpl) Validate(ctx context.Context, id AuthTokenID) (string, error) {
	if id == "" {
		return "", fmt.Errorf("invalid token id")
	}
	cr, err := as.storage.FindByAuthTokenID(ctx, id)
	if err != nil {
		return "", fmt.Errorf("token not found")
	}
	if cr.AuthTokenExpired() {
		return "", fmt.Errorf("token expired")
	}
	return cr.User, nil
}

func (as AuthServiceImpl) Expire(context.Context, string) error {
	return nil
}
