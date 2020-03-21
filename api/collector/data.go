package collector

import "context"

//Data struct
type Data struct {
	Time int64     `json:"time"`
	Data []float64 `json:"raw_data"`
}

type Repository interface {
	GetData(ctx context.Context) (Data, error)
}

func NilData() Data {
	return Data{
		Time: 0,
		Data: []float64{},
	}
}
