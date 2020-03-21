package main

import (
	"bio/datalistener"
	"bio/dbwriter"
	"bio/grpcSender"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	opcAddres                = "opc.tcp://127.0.0.99:48400"
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
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

func updateData() {
	var client = datalistener.GetClient(opcAddres)
	defer client.Close()
	var sess = dbwriter.GetSession()
	defer sess.Close()

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("178.130.44.169:9999", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	grpcClient := grpcSender.NewDataSendServiceClient(cc)

	for {
		var data, raw, _ = datalistener.GetData(client)
		dbwriter.PasteData(sess, data)
		request := &grpcSender.DataRequest{Data: convertTo32(raw)}
		_, err := grpcClient.SendData(context.Background(), request)
		print(err)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go updateData()
	for {
	}
}
