package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func addCount(date time.Time, count int) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	var existingCount int
	err = db.QueryRow("SELECT count FROM neo_count WHERE date=$1", date).Scan(&existingCount)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		_, err = db.Exec("INSERT INTO neo_count (date, count) VALUES ($1, $2)", date, count)
		if err != nil {
			return err
		}
	} else {
		_, err = db.Exec("UPDATE neo_count SET count=$1 WHERE date=$2", existingCount+count, date)
		if err != nil {
			return err
		}
	}

	return nil
}
