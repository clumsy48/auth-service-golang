package domain

import (
	"context"
	"fmt"
)

type OnboardingServiceImpl struct {
	storage Storage
	hasher  Hasher
}

func (o OnboardingServiceImpl) SignUp(ctx context.Context, credential User) error {
	if err := o.validate(credential); err != nil {
		return err
	}

	if _, err := o.storage.FindByEmail(ctx, credential.Email); err == nil {
		return fmt.Errorf("email already exists")
	}
	hash, err := o.hasher.Hash(credential.Password)
	if err == nil {
		credential.Password = hash
	}

	err = o.storage.SaveUser(ctx, credential)
	if err != nil {
		return err
	}
	return nil
}

func (o OnboardingServiceImpl) Delete(ctx context.Context, id AuthTokenID) error {
	return nil
}

func NewOnboardingService(storage Storage, hasher Hasher) *OnboardingServiceImpl {
	return &OnboardingServiceImpl{
		storage: storage,
		hasher:  hasher,
	}
}

func (o OnboardingServiceImpl) validate(user User) error {
	if user.FirstName == "" {
		return fmt.Errorf("missing first name")
	}

	if user.LastName == "" {
		return fmt.Errorf("missing last name")
	}

	if user.Email == "" {
		return fmt.Errorf("missing email name")
	}

	if user.Password == "" {
		return fmt.Errorf("missing password name")
	}
	return nil
}
