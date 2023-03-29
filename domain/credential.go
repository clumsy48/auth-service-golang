package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Credential struct {
	ID        CredentialID `json:"id"`
	AuthToken AuthToken    `json:"token"`
	User      string       `json:"user_id"`
}

type User struct {
	ID        CredentialID `json:"id"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
}

func (cr *Credential) GenerateAuthToken() error {
	h := sha256.New()
	if _, err := h.Write([]byte(fmt.Sprintf("%d-%s", time.Now().Unix(), randString()))); err != nil {
		return fmt.Errorf("sha256 write for token id failed %v", err)
	}
	cr.AuthToken = AuthToken{
		ID:        AuthTokenID(hex.EncodeToString(h.Sum(nil))),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
	return nil
}

func (cr *Credential) AuthTokenExpired() bool {
	return cr.AuthToken.ExpiresAt.Before(time.Now())
}

func randString() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 12)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
