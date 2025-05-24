package accounts

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type accountService struct {
	repository Respository
}

func NewService(r Respository) Service {
	return &accountService{r}
}

func (s *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {

	a := &Account{
		ID:   ksuid.New().String(),
		Name: name,
	}
	if err := s.repository.Save(ctx, *a); err != nil {
		return nil, err
	}

	return a, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {

	return s.repository.GetByID(ctx, id)
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {

	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}

	return s.repository.GetAll(ctx, skip, take)
}
