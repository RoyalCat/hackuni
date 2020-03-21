package collector

import (
	"context"
	"flag"

	"github.com/go-kit/kit/log"
	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(opcEnd string, logger log.Logger) Service {
	var endpoint = flag.String("endpoint", opcEnd, "OPC UA Endpoint URL")
	flag.Parse()

	ctx := context.Background()

	var c *opcua.Client = opcua.NewClient(*endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err := c.Connect(ctx); err != nil {
		logger.Log(err)
	}

	return &service{
		repo:   NewRepo(c, logger),
		logger: logger,
	}
}

func (s service) WaitData(ctx context.Context) (Data, error) {
	return s.repo.GetData(ctx)
}
