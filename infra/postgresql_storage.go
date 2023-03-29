package infra

import (
	"context"
	"gorm.io/gorm"
	"webserser/domain"
)

type postgresStorage struct {
	*gorm.DB
}

type credential struct {
	gorm.Model
	AuthToken domain.AuthToken `gorm:"embedded;embeddedPrefix:auth_token_"`
	User      string           `gorm:"user_id"`
}

type user struct {
	gorm.Model
	Email        string `gorm:"unique"`
	PasswordHash string
	FirstName    string
	LastName     string
}

func toCredential(a *credential) domain.Credential {
	return domain.Credential{
		ID:        domain.CredentialID(a.ID),
		AuthToken: a.AuthToken,
		User:      a.User,
	}
}

func fromCredential(a domain.Credential) *credential {
	return &credential{
		User:      a.User,
		AuthToken: a.AuthToken,
	}
}

func toUser(a *user) domain.User {
	return domain.User{
		ID:        domain.CredentialID(a.ID),
		Email:     a.Email,
		Password:  a.PasswordHash,
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}
}

func fromUser(a domain.User) *user {
	return &user{
		Email:        a.Email,
		PasswordHash: a.Password,
		FirstName:    a.FirstName,
		LastName:     a.LastName,
	}
}

func NewPostgresStorage(db *gorm.DB) (domain.Storage, error) {
	if err := db.AutoMigrate(&credential{}, &user{}); err != nil {
		return nil, err
	}

	return &postgresStorage{db}, nil
}

func (s postgresStorage) Save(ctx context.Context, a domain.Credential) error {
	row := fromCredential(a)
	if a.ID == domain.CredentialID(0) {
		return s.Create(row).Error
	}

	return s.WithContext(ctx).Model(&credential{}).Where("id = ?", a.ID).
		Updates(row).
		Error
}

func (s postgresStorage) SaveUser(ctx context.Context, a domain.User) error {
	row := fromUser(a)

	return s.WithContext(ctx).Model(&user{}).
		Save(row).
		Error
}

func (s postgresStorage) FindAll(ctx context.Context) ([]domain.Credential, error) {
	var rows []credential
	var credentials []domain.Credential

	tx := s.WithContext(ctx).Find(&rows)
	if tx.Error != nil {
		return credentials, tx.Error
	}

	for _, row := range rows {
		credentials = append(credentials, toCredential(&row))
	}

	return credentials, nil
}

func (s postgresStorage) FindByID(ctx context.Context, id domain.CredentialID) (domain.Credential, error) {
	var row credential

	tx := s.First(ctx, &row, id)
	if tx.Error != nil {
		return toCredential(&row), tx.Error
	}

	return toCredential(&row), nil
}

func (s postgresStorage) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var row user

	tx := s.WithContext(ctx).First(&row, "email = ?", email)
	if tx.Error != nil {
		return toUser(&row), tx.Error
	}

	return toUser(&row), nil
}

func (s postgresStorage) FindByAuthTokenID(ctx context.Context, id domain.AuthTokenID) (domain.Credential, error) {
	var row credential

	tx := s.WithContext(ctx).First(&row, "auth_token_id = ?", id)
	if tx.Error != nil {
		return toCredential(&row), tx.Error
	}

	return toCredential(&row), nil
}

func (s postgresStorage) DeleteByID(ctx context.Context, id domain.CredentialID) error {
	return s.WithContext(ctx).Delete(&credential{}, id).Error
}
