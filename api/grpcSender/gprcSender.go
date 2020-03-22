package grpcSender

import (
	context "context"
	"time"

	grpc "google.golang.org/grpc"
)

func convertTo32(ar []float64) []float32 {
	newar := make([]float32, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = float32(v)
	}
	return newar
}

func GetService(addres string) AnalystServiceClient {

	cc, err := grpc.Dial(addres, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		println("grpc error")
		return nil
	}

	return NewAnalystServiceClient(cc)
}

func SendData(client AnalystServiceClient, raw []float64) {
	request := &Enter{Message: convertTo32(raw)}
	client.Analyse(context.Background(), request)
}
