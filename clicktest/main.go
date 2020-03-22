package main

import (
	"log"
	"time"

	_ "github.com/mailru/go-clickhouse"
	"github.com/mailru/dbr"
)


func main() {
	connect, err := dbr.Open("clickhouse", "http://default:12345@78.140.223.19:8123", nil)
	if err != nil {
		log.Fatal(err)
	}
	var items []struct {
		CounterID uint32 `db:"CounterID"`
		Pressure float64 `db:"Pressure"`
		Humidity float64 `db:"Humidity"`
		TemperatureR float64 `db:"TemperatureR"`
		TemperatureA float64 `db:"TemperatureA"`
		pH float64 `db:"pH"`
		FlowRate float64 `db:"FlowRate"`
		CO float64 `db:"CO"`
		EventTime time.Time `db:"EventTime"`
	}
	sess := connect.NewSession(nil)
	query := sess.Select("*").From("test.data")
	if _, err := query.Load(&items); err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		print(item.Pressure)
	}
}