package main

import (
	"bio/datalistener"
	"bio/dbwriter"
	"bio/grpcSender"
	"time"
)

const (
	opcAddres                = "opc.tcp://127.0.0.99:48400"
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
	grpcAddres               = "localhost:9999"
)

func updateData() {
	var client = datalistener.GetClient(opcAddres)
	defer client.Close()
	var sess = dbwriter.GetSession()
	defer sess.Close()
	var grpcService = grpcSender.GetService(grpcAddres)

	for {
		var data, raw, _ = datalistener.GetData(client)
		raw = append(raw, float64(data.EventTime.Unix()))

		grpcSender.SendData(grpcService, raw)
		dbwriter.PasteData(sess, data)

		time.Sleep(1 * time.Second)
	}
}

func main() {
	go updateData()
	for {
	}
}
