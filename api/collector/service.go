package collector

import "context"

type Service interface {
	WaitData(ctx context.Context) (Data, error)
}
