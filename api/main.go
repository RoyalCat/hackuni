package main

import (
	"bio/datalistener"
	"bio/dbwriter"
	"bio/grpcSender"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const configfile = "config.json"

type Configuration struct {
	OpcAddres  string `json:"OpcAddres"`
	GrpcAddres string `json:"GrpcAddres"`
	DbAddres   string `json:"DbAddres"`
}

func loadconfig(path string) Configuration {
	config := Configuration{}

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal("Config error ", err.Error())
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Config error ", err.Error())
	}
	if config.DbAddres == "" || config.GrpcAddres == "" || config.OpcAddres == "" {
		log.Fatal("Config value error")
	}

	return config
}

func servicesUpdate(dbsess *sql.DB, grpcCleint *grpc.ClientConn, config Configuration) {
	dbsess = dbwriter.GetSession(config.DbAddres)
	grpcCleint = grpcSender.GetService(config.GrpcAddres)
}

func updateData(config Configuration) {
	println("Connecting to opc server...")
	var client = datalistener.GetClient(config.OpcAddres)
	defer client.Close()
	println("Connected")

	println("Connecting to database...")
	var sess = dbwriter.GetSession(config.DbAddres)
	defer sess.Close()
	println("Connected")

	println("Connecting to analyser(grpc)...")
	var grpcClient = grpcSender.GetService(config.GrpcAddres)
	var grpcService grpcSender.AnalystServiceClient
	if grpcClient != nil {
		grpcService = grpcSender.NewAnalystServiceClient(grpcClient)
	}

	println("Connected")

	println("Running...")
	for {
		var data datalistener.Item
		var raw []float64
		if client != nil {
			data, raw, _ = datalistener.GetData(client)
			raw = append(raw, float64(data.EventTime.Unix()))
		} else {
			data = datalistener.NilData()
			raw = nil
		}
		if data != datalistener.NilData() {
			if grpcService != nil {
				grpcSender.SendData(grpcService, raw)
			}
			if sess != nil {
				var err = dbwriter.PasteData(sess, data)
				if err != nil {
					sess = nil
				}
			}
		}
		if grpcClient != nil && sess != nil && client != nil {
			time.Sleep(1 * time.Second)
		}

		if sess == nil {
			sess = dbwriter.GetSession(config.DbAddres)
			if grpcClient != nil {
				time.Sleep(1)
			}
		}

		if client == nil {
			client = datalistener.GetClient(config.OpcAddres)
			if grpcClient != nil {
				time.Sleep(1)
			}
		}

		if grpcClient == nil {
			grpcClient = grpcSender.GetService(config.GrpcAddres)
			if grpcClient != nil {
				grpcService = grpcSender.NewAnalystServiceClient(grpcClient)
			}
		}

		println(time.Now().Format("15:04:05"), "tick")
	}
}

func main() {
	println("Initialization...")
	config := loadconfig(configfile)
	updateData(config)
}
