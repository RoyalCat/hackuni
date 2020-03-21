package collector

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type waitDataRequest struct {
}

type waitDataResponse struct {
	Data string `json:"data"`
	Err  error  `json:"error"`
}

func (r waitDataResponse) error() error { return r.Err }

func makeBookCargoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(waitDataRequest)
		data, err := s.WaitData(ctx)
		return waitDataResponse{Data: data, Err: err}, nil
	}
}