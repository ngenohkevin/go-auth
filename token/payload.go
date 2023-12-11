package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrExpiredTokens = errors.New("token has expired")
	ErrInvalidToken  = errors.New("token is invalid")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredTokens
	}
	return nil
}
