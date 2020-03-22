package main

import (
	"bio/datalistener"
	"bio/dbwriter"
	"bio/grpcSender"
	"encoding/json"
	"log"
	"os"
	"time"
)

const configfile = "config.json"

type Configuration struct {
	OpcAddres  string `json:"opcAddres"`
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

func updateData(config Configuration) {
	var client = datalistener.GetClient(config.OpcAddres)
	defer client.Close()
	var sess = dbwriter.GetSession(config.DbAddres)
	defer sess.Close()
	var grpcService = grpcSender.GetService(config.GrpcAddres)

	for {
		var data, raw, _ = datalistener.GetData(client)
		raw = append(raw, float64(data.EventTime.Unix()))

		grpcSender.SendData(grpcService, raw)
		dbwriter.PasteData(sess, data)

		time.Sleep(1 * time.Second)
	}
}

func main() {
	config := loadconfig(configfile)
	updateData(config)
}
