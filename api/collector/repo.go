package collector

import (
	"context"
	"errors"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

const minID = 2
const maxID = 8

type repo struct {
	client *opcua.Client
	logger log.Logger
}

func NewRepo(client *opcua.Client, logger log.Logger) Repository {
	return &repo{
		client: client,
		logger: log.With(logger, "repo", "opc"),
	}
}

func (repo *repo) GetData(ctx context.Context) (Data, error) {
	req := ua.ReadRequest{
		MaxAge:             2000,
		TimestampsToReturn: ua.TimestampsToReturnBoth,
		NodesToRead:        []*ua.ReadValueID{},
	}

	for i := uint32(minID); i <= maxID; i++ {
		var nodeid = ua.NewNumericNodeID(2, i)
		req.NodesToRead = append(req.NodesToRead, &ua.ReadValueID{NodeID: nodeid})
	}

	resp, err := repo.client.Read(&req)
	if err != nil {
		repo.logger.Log("Read failed: %s", err)
		return NilData(), err
	}

	if resp.Results[0].Status != ua.StatusOK {
		repo.logger.Log("Status not OK: %v", resp.Results[0].Status)
		return NilData(), err
	}

	var out []float64

	for _, res := range resp.Results {
		out = append(out, float64(res.Value.Int())+res.Value.Float())
	}

	return Data{
		Data: out,
		Time: resp.Header().Timestamp.Unix(),
	}, nil
}
