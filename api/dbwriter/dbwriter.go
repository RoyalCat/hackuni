package dbwriter

import (
	"bio/datalistener"
	"database/sql"
	"log"

	_ "github.com/mailru/go-clickhouse"
)

func GetSession(addres string) *sql.DB {
	connect, err := sql.Open("clickhouse", addres)
	if err != nil {
		print("db error")
		//log.Fatal(err)
		return nil
	}
	return connect
}
func PasteData(conn *sql.DB, data datalistener.Item) {
	tx, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(`
		INSERT INTO test.data (
			Pressure,
			Humidity,
			TemperatureR,
			TemperatureA,
			pH,
			FlowRate,
			CO,
			EventTime
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?
		)`)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := stmt.Exec(
		data.Pressure,
		data.Humidity,
		data.TemperatureA,
		data.TemperatureR,
		data.PH,
		data.FlowRate,
		data.CO,
		data.EventTime.Unix(),
	); err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
