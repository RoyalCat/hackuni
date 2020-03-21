package main

import (
	"bio/collector"
	"context"

	"github.com/go-kit/kit/log"
)

const (
	opcAddres                = "opc.tcp://127.0.0.99:48400"
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
)

func main() {
	var ctx = context.Background()

	var logr = log.NewNopLogger()
	var col = collector.NewService(opcAddres, logr)
	var res, _ = col.WaitData(ctx)
	print(res.Data)
}
