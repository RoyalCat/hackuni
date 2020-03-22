package grpcSender

import (
	context "context"
	"log"

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
	var opts = grpc.WithInsecure()

	cc, err := grpc.Dial(addres, opts)
	if err != nil {
		log.Fatal(err)
	}

	return NewAnalystServiceClient(cc)
}

func SendData(client AnalystServiceClient, raw []float64) {
	request := &Enter{Message: convertTo32(raw)}
	client.Analyse(context.Background(), request)
}
