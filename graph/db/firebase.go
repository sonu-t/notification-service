package db

import (
	"context"
	"errors"
)

type FirebaseRepo interface {
	StoreToken(ctx context.Context, userId int, token string) error
	GetToken(ctx context.Context, userId int) (string, error)
}

type firebaseRepo struct {
	tokens map[int]string
}

func (f firebaseRepo) StoreToken(ctx context.Context, userId int, token string) error {
	f.tokens[userId] = token
	return nil
}

func (f firebaseRepo) GetToken(ctx context.Context, userId int) (string, error) {
	if token, ok := f.tokens[userId]; ok {
		return token, nil
	}
	return "", errors.New("no token registered")
}

func NewFirebaseRepo() FirebaseRepo {
	return &firebaseRepo{tokens: map[int]string{}}
}
