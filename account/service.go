package account

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	CreateAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type accountService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &accountService{
		repo: repo,
	}
}

func (s *accountService) CreateAccount(ctx context.Context, name string) (*Account, error) {
	account := &Account{
		ID:   ksuid.New().String(),
		Name: name,
	}
	if err := s.repo.CreateAccount(ctx, *account); err != nil {
		return nil, err
	}

	return account, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	return s.repo.GetAccountByID(ctx, id)
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}

	return s.repo.GetAccounts(ctx, skip, take)
}
