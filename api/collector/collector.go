package collector

import "context"

//Collector interface
type Collector interface {
	waitData(ctx context.Context) (Data, error)
}
