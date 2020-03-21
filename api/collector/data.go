package collector

import "context"

//Data struct
type Data struct {
	time int64  `json:"time"`
	data string `json:"raw_data"`
}

type Repository interface {
	GetData(ctx context.Context) (string, error)
}
