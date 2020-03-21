package collector

import (
	"context"

	"github.com/go-kit/kit/log"
)

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s service) WaitData(ctx context.Context) (string, error) {
	return "Succ", nil
}
