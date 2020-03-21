package main

import (
	"bio/collector"
	"context"
	"flag"
	"log"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

const (
	opcAddres                = "opc.tcp://127.0.0.99:48400"
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
)

func main() {
	var (
		endpoint = flag.String("endpoint", "opc.tcp://localhost:4840", "OPC UA Endpoint URL")
		nodeID   = flag.String("node", "", "NodeID to read")
	)
	flag.Parse()

	ctx := context.Background()

	var c opcua.Client = opcua.NewClient(*endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	var col collector.Service
	col = collector.NewService(c, nil)

}
